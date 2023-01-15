document.getElementById("_submit").addEventListener("click", function () {
    var _title = document.getElementById("_title").value
    var _url = document.getElementById("_url").value
    const data = {
        title: _title,
        url: _url,
    };

    fetch('http://127.0.0.1:8569/register', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Accept': 'application/json',
            'Content-type': 'application/json; charset=UTF-8'
        }
    })
        .then(response => response.json())
        .then(json => {
            document.location = "http://127.0.0.1:8569/result"
        });
})
