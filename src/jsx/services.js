import fetch from 'isomorphic-fetch';


class Services {
    listStandardDeviations() {
        return fetch('/standardDeviation').then(sds => sds.json());
    }

    saveStandardDeviation(numbers) {
        var request = new Request('/standardDeviation', {
            method: 'POST',
            mode: 'cors',
            body: numbers
        });
        return fetch(request).then(response => response.json());
    }
}

module.exports = new Services()
