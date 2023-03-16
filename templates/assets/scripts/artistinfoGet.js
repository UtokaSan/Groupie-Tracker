const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get('artist');
const idAlbum = urlParams.get('name');
const data = {id}
const dataAlbum = {idAlbum}
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
            let textNode = document.createTextNode(`Members : ${data.artist.members}`)
            createParagraph.appendChild(textNode)
            document.querySelector(".test").appendChild(createParagraph)
            document.querySelector(".title_paralax").textContent = data.artist.name
            document.querySelector(".about__p").textContent = `Creation Date : ${data.artist.creationDate}`

            for (var i = 0; i < data.dates.dates.length; i++) {
                const servList = document.querySelector(".serv__list");

                const servItem = document.createElement("div");
                servItem.classList.add("serv__item");

                const servItemArrow = document.createElement("span");
                servItemArrow.classList.add("serv__item-arrow");
                servItemArrow.setAttribute("data-speed", "800");

                const arrowImg = document.createElement("img");
                arrowImg.setAttribute("src", "assets/media/arrow.svg");
                arrowImg.setAttribute("alt", "");

                servItemArrow.appendChild(arrowImg);

                servItem.appendChild(servItemArrow);

                const servItemTxt = document.createElement("div");
                servItemTxt.classList.add("serv__item-txt");
                servItemTxt.textContent = data.dates.dates[i]

                servItem.appendChild(servItemTxt);
                servList.appendChild(servItem);
            }

            const map = new google.maps.Map(document.getElementById("map"), {
                center: {lat: 48.8566, lng: 2.3522},
                zoom: 2
            });
            data.location.locations.forEach(city => {
                const geocoder = new google.maps.Geocoder();
                geocoder.geocode({ address: city }, (results, status) => {
                    if (status === 'OK') {
                        const lat = results[0].geometry.location.lat();
                        const lng = results[0].geometry.location.lng();
                        new google.maps.Marker({
                            position: { lat, lng },
                            map,
                            title: city
                        });
                    } else {
                        console.log("Geocode n'a pas fonctionné :" + status);
                    }
                });
            });
        })
        .catch(error => console.error(error))
})
document.addEventListener('DOMContentLoaded', function () {
    fetch('/get/artistinfoalbum', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(dataAlbum)
    })
        .then(response => response.json())
        .then(data => {
            console.log(data)
            let createParagraph = document.createElement("p");
            let createParagraph1 = document.createElement("p");
            createParagraph.className = "about__p"
            createParagraph1.className = "about__p"
            let textNode = document.createTextNode(`Nombre d'écoute : ${data.allListeners.artist.stats.listeners}`)
            str = data.allListeners.artist.bio.summary.replace(/<a[^>]*>([^<]*)<\/a>/g, "$1");
            str = str.replace("Read more on Last.fm", '');
            let textNode1 = document.createTextNode(`Autobiography : ${str}`)
            createParagraph.appendChild(textNode)
            createParagraph1.appendChild(textNode1)
            document.querySelector(".test").appendChild(createParagraph)
            document.querySelector(".test").appendChild(createParagraph1)
            for (let i = 0, j = 4; i <= 4; i++, j++) {
                if (data.allAlbum.topalbums.album[i]["name"] === "(null)") {
                    console.log("No Image")
                } else {
                    let imageTake = document.querySelectorAll(".work__item-img");
                    let parentImage = document.querySelector(".work__item");
                    let targetElement = imageTake[i];
                    let parentImageTarget = parentImage[i]
                    let createImage = document.createElement("img");
                    createImage.setAttribute("src", data.allAlbum.topalbums.album[j].image[3]['#text']);
                    createImage.onerror = function() {
                        parentImageTarget.style.display = 'none';
                    };
                    targetElement.appendChild(createImage);
                    console.log(data.allAlbum.topalbums.album[j].image[3]['#text']);
                }
            }
        })
        .catch(error => console.error(error))
})