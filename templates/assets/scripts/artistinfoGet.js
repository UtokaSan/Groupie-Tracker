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
            let createParagraph = document.createElement("p");
            createParagraph.className = "about__p"
            let textNode = document.createTextNode(data.artist.members)
            createParagraph.appendChild(textNode)
            document.querySelector(".test").appendChild(createParagraph)
            document.querySelector(".title_paralax").textContent = data.artist.name
            document.querySelector(".about__p").textContent = data.artist.creationDate
            for (var i = 0; i < data.dates.dates.length; i++) {
                let createSpan = document.createElement("span")
                createSpan.className = "serv__item-text"
                let textNode = document.createTextNode(data.dates.dates[i])
                createSpan.appendChild(textNode)
                document.querySelector(".serv__item-txt").appendChild(createSpan)
            }
        })
        .catch(error => console.error(error))
})