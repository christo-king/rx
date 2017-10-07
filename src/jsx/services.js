import fetch from 'isomorphic-fetch';


class Services {
    listStandardDeviations() {
        return fetch('/standardDeviation').then(sds => sds.json());
    }

    saveStandardDeviation(numbers) {
        let body = JSON.stringify({points: numbers});
        var request = new Request('/standardDeviation', {
            method: 'POST',
            mode: 'cors',
            body: body
        });
        return fetch(request).then(response => this.shapeResponse(response));
    }

    shapeResponse(response) {
        if (response.status >= 400) {
            throw "Server error";
        } else {
            return response.json();
        }
    }
}

module.exports = new Services()
