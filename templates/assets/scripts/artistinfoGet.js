const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get('artist');
const data = {id}
document.addEventListener('DOMContentLoaded', function () {
    fetch('/get/artistinfo', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            document.querySelector(".title_paralax").texteContent = data.artist.name
            document.querySelector(".about__p").textContent = data.artist.creationDate
        })
        .catch(error => console.error(error))
})