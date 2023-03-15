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
