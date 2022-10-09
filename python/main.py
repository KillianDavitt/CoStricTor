import hashlib
import numpy as np
import pandas as pd


class Filter:
    def __init__(self, size, threshold):
        self.size = size
        self.filter = np.zeros(size, dtype='int32')
        self.threshold = threshold
        self.count = 0

    def add(self, hashmap):
        self.filter += hashmap
        self.count += 1

    def check(self, array):
        if self.count == 0:
            return False
        values = self.filter[np.nonzero(array)] / self.count > self.threshold
        return np.all(values)


# create some data from various distributions
def set_of_sites(num_sites=100, proportion_https=0.7, proportion_hsts_of_https=0.3):
    # Generates a set of strings for simulation purposes.
    sites = pd.read_csv('./majestic_million.csv', nrows=num_sites, usecols=['Domain'])
    sites['https'] = np.random.binomial(1, proportion_https, size=sites.shape[0])
    sites['hsts'] = sites['https'] & np.random.binomial(1, proportion_hsts_of_https, size=sites.shape[0])
    return sites


def get_sample_probabilities(num_sites=100, decay="Measured", background=None):
    probs = np.zeros(num_sites)
    if decay == "Measured":
        df = pd.read_csv('./majestic_million.csv', nrows=num_sites, usecols=['RefSubNets'])
        temp = np.array(df['RefSubNets'])
        probs = temp / temp.sum()
    elif decay == "Zipf":
        temp = - 1.13 / np.arange(1, num_sites + 1)
        probs = temp / temp.sum()
    elif decay == "Linear":
        temp = np.arrange(num_sites, 0, -1)
        probs = temp / temp.sum()
    elif decay == "Constant":
        probs = np.full(num_sites, 1 / num_sites)
    elif decay == "Exponential":
        temp = -np.array(range(1, num_sites + 1))
        temp = np.exp(temp) + background
        temp = temp / temp.sum()
        probs = temp
    else:
        raise Exception('decay must be one of ["Linear", "Exponential", "Constant", "Measured", "Zipf"]')
    return probs


def sample_sites(N, sites, probs):
    return np.random.choice(a=sites['Domain'], size=N, replace=True, p=probs)


def create_hashmap(size, site, hash_count):
    # todo: cohorts
    hashmap = np.zeros(size, dtype='int32')
    encoded = site.encode()
    for i in range(0, hash_count):
        h = hashlib.md5(str(i).encode())
        h.update(encoded)
        hashmap[int(h.hexdigest(), 16) % size] += 1
    return hashmap


def main():
    # todo: read from command line or csv
    num_of_samples = 100000
    num_sites = 10000
    proportion_https = 0.7
    proportion_hsts_of_https = 0.3
    cohorts = 2 # todo: use this?
    size_of_filter = 1024
    decay = "Measured"
    hash_count = 2
    threshold = 0.02

    sites = set_of_sites(num_sites, proportion_https, proportion_hsts_of_https)
    probs = get_sample_probabilities(num_sites, decay)
    samples = sample_sites(num_of_samples, sites, probs)
    sites['hashmap'] = sites.apply(lambda r: create_hashmap(size_of_filter, r['Domain'], hash_count), axis=1)

    hsts_filter = Filter(size_of_filter, threshold)
    nohttps_filter = Filter(size_of_filter, threshold)

    for sample in samples:
        site = sites[sites['Domain'] == sample]
        if site['hsts'].values[0] == 1:
            hsts_filter.add(site['hashmap'].values[0])
        elif site['https'].values[0] == 0:
            nohttps_filter.add(site['hashmap'].values[0])

    count_hsts = 0
    count_benefit = 0
    count_no_benefit = 0

    count_insecure = 0
    count_disaster = 0
    count_no_disaster = 0

    for row in sites.iterrows():
        hsts_result = hsts_filter.check(row[1]['hashmap'])
        nohttps_result = nohttps_filter.check(row[1]['hashmap'])

        if row[1]['hsts']:
            count_hsts += 1
            if hsts_result:
                if nohttps_result:
                    count_no_benefit += 1
                else:
                    count_benefit += 1
        if row[1]['https'] == 0:
            count_insecure += 1
            if hsts_result:
                if nohttps_result:
                    count_no_disaster += 1
                else:
                    count_disaster += 1

    print('hsts', count_hsts, count_benefit, count_no_benefit)
    print('nohttps', count_insecure, count_disaster, count_no_disaster)


if __name__ == '__main__':
    main()
