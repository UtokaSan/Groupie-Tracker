document.addEventListener('DOMContentLoaded', function () {
    fetch('/post/searchbar')
        .then(response => response.json())
        .then(data => {
            let valueSearch = document.getElementById("search-input");
            let suggestions = document.getElementById("suggestions");
            let form = document.querySelector('.search-form');
            valueSearch.addEventListener("input", function(event) {
                suggestions.innerHTML = "";
                if (valueSearch.value.length > 3) {
                    for (let artist of data.artists) {
                        if (artist.name.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                            let option = document.createElement("option");
                            option.value = `${artist.name} -Artist`;
                            suggestions.appendChild(option);
                        }
                        for (let member of artist.members) {
                            if (member.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                                let option = document.createElement("option");
                                option.value = `${member} -Member`;
                                suggestions.appendChild(option);
                            }
                        }
                        if (artist.creationDate.toString().startsWith(valueSearch.value.toLowerCase())) {
                            let option = document.createElement("option");
                            option.value = `${artist.creationDate} - ${artist.name} -Creation Date`;
                            suggestions.appendChild(option);
                        }

                        for (let index of data.location.index) {
                            for (let location of index.locations) {
                                location = location.replace(/[-_]/g, ' ');
                                if (location.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                                    let artist = data.artists.find(function (artist) {
                                        return artist.id === index.id;
                                    });
                                    let option = document.createElement("option");
                                    option.value = `${location} - ${artist.name}`
                                    if (!suggestions.querySelector(`option[value="${option.value}"]`)) {
                                        suggestions.appendChild(option);
                                    }
                                }
                            }
                        }
                    }
                }
            });

            form.addEventListener('submit', function(event) {
                event.preventDefault();
                data.artists.forEach(function (artist) {
                    if (valueSearch.value === artist.name + " -Artist") {
                        window.location.href = `http://localhost:8080/artistinfo?artist=${artist.id}&${artist.name}`;
                    }
                    artist.members.forEach(function (member) {
                        if (valueSearch.value === member + " -Member") {
                            window.location.href = `http://localhost:8080/artistinfo?artist=${artist.id}&${artist.name}`;
                        }
                    });
                    if (valueSearch.value === `${artist.creationDate} - ${artist.name} -Creation Date`) {
                        window.location.href = `http://localhost:8080/artistinfo?artist=${artist.id}&${artist.name}`;
                    }
                    data.location.index.forEach(function (index) {
                        index.locations.forEach(function (location) {
                            location = location.replace(/[-_]/g, ' ');
                            if (valueSearch.value === `${location} - ${artist.name}`) {
                                window.location.href = `http://localhost:8080/artistinfo?artist=${artist.id}&${artist.name}`;
                            }
                        })
                    })
                });
            });
        })
        .catch(error => console.error(error))
});
