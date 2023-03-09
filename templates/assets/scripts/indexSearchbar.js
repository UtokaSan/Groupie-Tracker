document.addEventListener('DOMContentLoaded', function () {
    fetch('/post/searchbar')
        .then(response => response.json())
        .then(data => {
            let valueSearch = document.getElementById("search-input")
            let button = document.querySelector(".search-button");
            /*console.log(data.relation.index[0].datesLocations);*/
            console.log(data.location.index)
            console.log(data.artists[0].creationDate)
            valueSearch.addEventListener("input", function(event) {
                let suggestions = document.getElementById("suggestions");
                suggestions.innerHTML = "";
                data.artists.forEach(function (artist) {
                    if (artist.name.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                        let option = document.createElement("option");
                        option.value = artist.name + " -Artist";
                        suggestions.appendChild(option);
                    }
                    artist.members.forEach(function (member) {
                        if (member.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                            let option = document.createElement("option");
                            option.value = member + " -Member";
                            suggestions.appendChild(option);
                        }
                    });
                    if (artist.creationDate.toString().startsWith(valueSearch.value.toLowerCase())) {
                        let option = document.createElement("option");
                        option.value = artist.creationDate + " -Creation Date";
                        suggestions.appendChild(option);
                    }
                });
                data.location.index.forEach(function (index) {
                    index.locations.forEach(function (location) {
                        if (location.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                            let option = location + " -Location";
                            valueSearch.setAttribute("placeholder", option)
                        }
                    });
                });

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
                    data.location.index.forEach(function (location) {
                        if (event.target.value === location + " -Location") {
                            window.location.href = `http://localhost:8080/artistinfo?artist=${location.id}`;
                        }
                    })
                });
            })


            /*           valueSearch.addEventListener("input", function(event) {
                           let suggestions = document.getElementById("suggestions");
                           suggestions.innerHTML = "";
                           data.relation.index.forEach(function (value, index) {
                               Object.keys(value.datesLocations).forEach(function(key) {
                                   if (key.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                                       let option = document.createElement("option");
                                       option.value = value.datesLocations[key] + " -Relation";
                                       suggestions.appendChild(option);
                                   }
                               });
                           });
                       });
                       valueSearch.addEventListener("input", function(event) {
                           data.relation.index.forEach(function (value, index) {
                               Object.keys(value.datesLocations).forEach(function(key) {
                                   if (event.target.value === value.datesLocations[key] + " -Relation") {
                                       window.location.href = `http://localhost:8080/artistinfo?artist=${value.id}`;
                                   }
                               });
                           });
                       });*/
        })
        .catch(error => console.error(error))
})