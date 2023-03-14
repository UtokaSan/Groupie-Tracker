const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get('artist');
const data = {id}
console.log("test")
document.addEventListener('DOMContentLoaded', function () {
    fetch('/get/artistinfo', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            document.querySelector(".title_paralax").textContent = data.artist.name

            document.querySelector(".about__p").textContent = data.artist.creationDate
            console.log(data)
        })
        .catch(error => console.error(error))
})
