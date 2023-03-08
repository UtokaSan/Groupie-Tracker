
const urlParams = new URLSearchParams(window.location.search);
const test = document.getElementById("test");
const id = urlParams.get('genre');

const data = {id};
document.addEventListener('DOMContentLoaded', function () {
    fetch('/api/genre', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            data.forEach(function callback(value, index)  {
                var button = document.createElement("button");
                const img = document.createElement('img');
                const div = document.createElement('div');
                button.classList.add("name" + index)
                button.textContent = value.name;
                img.src = value.image;
                if (index >= 6 && index < 12) {
                    div.classList.add("artist1");
                } else if (index >= 12 && index < 18) {
                    div.classList.add("artist2");
                } else if (index >= 18 && index < 24) {
                    div.classList.add("artist2");
                } else if (index >= 24) {
                    div.classList.add("artist3");
                } else {
                    div.classList.add("artist");
                }
                div.appendChild(button);
                div.appendChild(img);
                test.appendChild(div);
                button.addEventListener('click', function () {
                    let urlParameter = new URLSearchParams();
                    urlParameter.append('artist', value.id);
                    let urlNewPage = `http://localhost:8080/artistinfo?${urlParameter.toString()}`
                    window.location.href = urlNewPage
                });
            });
        })
        .catch(error => console.error(error))
})