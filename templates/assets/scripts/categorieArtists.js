const urlParamsGetGenre = new URLSearchParams(window.location.search);
const test = document.getElementById("test");
const id = urlParamsGetGenre.get('genre');

const data = {id};


function filterArtists() {
    const dateCreation = document.getElementById("dateCreation").value;
    const members = document.getElementById("members").value;
    const firstAlbum = document.getElementById("album-date").value;
    const locationArtist = document.getElementById("location").value;
    const artists = document.querySelectorAll('.artist');
    const artist1 = document.querySelectorAll('.artist1');
    const artist2 = document.querySelectorAll('.artist2');

    artists.forEach(artist => {
        const years = artist.querySelector('button').getAttribute('data-years');
        const nbMembers = artist.querySelector('button').getAttribute('data-members');
        const album = artist.querySelector('button').getAttribute('data-years-first-album');
        const location = artist.querySelector('button').getAttribute('data-location');

        if ((dateCreation && years < dateCreation) ||
            (members && nbMembers !== members) ||
            (firstAlbum && album > firstAlbum) ||
            (locationArtist && location === locationArtist)) {
            artist.style.display = "none";
        } else {
            if (!locationArtist || location.split(',').includes(locationArtist) || locationArtist === "location") {
                artist.style.display = "block";
            } else {
                artist.style.display = "none";
            }
        }
    });

    artist1.forEach(artist => {
        const years = artist.querySelector('button').getAttribute('data-years');
        const nbMembers = artist.querySelector('button').getAttribute('data-members');
        const album = artist.querySelector('button').getAttribute('data-years-first-album');
        const location = artist.querySelector('button').getAttribute('data-location');

        if ((dateCreation && years < dateCreation) ||
            (members && nbMembers !== members) ||
            (firstAlbum && album > firstAlbum) ||
            (locationArtist && location === locationArtist)) {
            artist.style.display = "none";
        } else {
            if (!locationArtist || location.split(',').includes(locationArtist) || locationArtist === "location") {
                artist.style.display = "block";
            } else {
                artist.style.display = "none";
            }
        }
    });

    artist2.forEach(artist => {
        const years = artist.querySelector('button').getAttribute('data-years');
        const nbMembers = artist.querySelector('button').getAttribute('data-members');
        const album = artist.querySelector('button').getAttribute('data-years-first-album');
        const location = artist.querySelector('button').getAttribute('data-location');

        if ((dateCreation && years < dateCreation) ||
            (members && nbMembers !== members) ||
            (firstAlbum && album > firstAlbum) ||
            (locationArtist && location === locationArtist)) {
            artist.style.display = "none";
        } else {
            if (!locationArtist || location.split(',').includes(locationArtist) || locationArtist === "location") {
                artist.style.display = "block";
            } else {
                artist.style.display = "none";
            }
        }
    });
}

document.addEventListener('DOMContentLoaded', function () {
    fetch('/api/genre', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            data.artists.forEach(function callback(artist, index)  {
                var button = document.createElement("button");
                const img = document.createElement('img');
                const div = document.createElement('div');
                img.src = artist.image;
                button.classList.add("name" + index)
                button.textContent = artist.name;
                button.setAttribute("data-years", artist.creationDate)
                button.setAttribute("data-years-first-album", artist.firstAlbum)
                button.setAttribute("data-members", artist.members.length)
                if (index >= 6 && index < 12) {
                    div.classList.add("artist1");
                } else if (index >= 12 && index < 18) {
                    div.classList.add("artist2");
                } else {
                    div.classList.add("artist");
                }
                div.appendChild(img);
                div.appendChild(button);
                test.appendChild(div);
                data.location.index.forEach(function callback(value, index) {
                    if (value.id === artist.id) {
                        let optionCreate = document.createElement("option");
                        //Corriger car il ne met pas toutes locations dans les options
                        value.locations.forEach(function (loc) {
                            optionCreate.textContent = loc
                            document.querySelector("#location").appendChild(optionCreate)
                        })
                        button.setAttribute("data-location", value.locations)
                    }
                })
                button.addEventListener('click', function () {
                    let urlParameter = new URLSearchParams();
                    let urlParameterName = new URLSearchParams()
                    urlParameter.append('artist', artist.id);
                    urlParameterName.append('name', artist.name)
                    let urlNewPage = `http://localhost:8080/artistinfo?${urlParameter.toString()}&${urlParameterName.toString()}`
                    window.location.href = urlNewPage
                });
            });
            filterArtists();
        })
        .catch(error => console.error(error))
})
