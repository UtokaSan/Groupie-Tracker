
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
                button.classList.add("name" + index)
                button.textContent = value.name;
                img.src = value.image;
                test.appendChild(button);
                test.appendChild(img);
            });
        })
        .catch(error => console.error(error))
})

window.addEventListener("load", function(){
    setTimeout(function(){
        document.querySelector(".progress-bar").style.width = "0%";
    }, 10000);
});