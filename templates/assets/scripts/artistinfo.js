
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
                button.classList.add("name" + index)
                button.textContent = value.name;
                test.appendChild(button);
            });
        })
        .catch(error => console.error(error))
})
metalbutton.addEventListener('click', function() {
    let urlParameterMetal = new URLSearchParams();
    urlParameterMetal.append('genre', 'metal');

    let urlNewPageMetal = `http://localhost:8080/categorie?${urlParameterMetal.toString()}`
    window.location.href = urlNewPageMetal
});
setTimeout(function() {
    window.location.href = "http://localhost:3001/";
}, 10000);

function annulerRedirection() {
    clearTimeout();
}