let dataArtist = [];
document.addEventListener('DOMContentLoaded', function () {
    fetch('/post/searchbar')
        .then(response => response.json())
        .then(data => {
            let valueSearch = document.getElementById("search-input")
            let button = document.querySelector(".search-button");
            let takeFirstLetter = valueSearch.value.substring(0, 3)
            /*console.log(data.relation.index[0].datesLocations);*/
            valueSearch.addEventListener("input", function(event) {
                let suggestions = document.getElementById("suggestions");
                suggestions.innerHTML = "";
                if (valueSearch.value.length > 3) {
                    data.artists.forEach(function (artist) {
                        if (artist.name.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                            let option = document.createElement("option");
                            option.value = artist.name + " -Artist";
                            suggestions.appendChild(option);
                        }
                        artist.members.forEach(function (member) {
                            if (member.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                                let option = document.createElement("option");
                                suggestions.appendChild(option);
                                option.value = member + " -Member";
                            }
                        });
                        if (artist.creationDate.toString().startsWith(valueSearch.value.toLowerCase())) {
                            let option = document.createElement("option");
                            option.value = artist.creationDate + " - " + artist.name + " -Creation Date";
                            suggestions.appendChild(option);
                        }
                        data.location.index.forEach(function (index) {
                            index.locations.forEach(function (location) {
                                if (location.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                                    let option = document.createElement("option");
                                    option.value = artistName;
                                    suggestions.appendChild(option);
                                }
                            });
                        });
                    });
                }
                valueSearch.addEventListener("input", function (event) {
                    data.artists.forEach(function (artist) {
                        if (event.target.value === artist.name + " -Artist") {
                            window.location.href = `http://localhost:8080/artistinfo?artist=${artist.id}`;
                        }
                        artist.members.forEach(function (member) {
                            if (event.target.value === member + " -Member") {
                                window.location.href = `http://localhost:8080/artistinfo?artist=${artist.id}`;
                            }
                        });
                        if (event.target.value === artist.creationDate.toString() + " -Creation Date") {
                                window.location.href = `http://localhost:8080/artistinfo?artist=${artist.id}`;
                        }
                    });
                    data.location.index.forEach(function (index) {
                        index.locations.forEach(function (location) {
                            if (event.target.value === location + " -Artist") {
                                window.location.href = `http://localhost:8080/artistinfo?artist=${index.id}`;
                            }
                        })
                    })
                });
            })
        })
        .catch(error => console.error(error))
})