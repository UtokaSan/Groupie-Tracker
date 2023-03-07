document.addEventListener('DOMContentLoaded', function () {
    fetch('/post/searchbar')
        .then(response => response.json())
        .then(data => {
            let valueSearch = document.getElementById("search-input")
            let button = document.querySelector(".search-button");
           /* console.log(data.relation.datesLocations[0])*/
            valueSearch.addEventListener("input", function(event) {
                let suggestions = document.getElementById("suggestions");
                suggestions.innerHTML = "";
                data.artists.forEach(function (value, index) {
                    if (value.name.toLowerCase().startsWith(valueSearch.value.toLowerCase())) {
                        let option = document.createElement("option");
                        option.value = value.name + " -Artist";
                        suggestions.appendChild(option);
                    }
                });
            });
            valueSearch.addEventListener("input", function(event) {
                data.artists.forEach(function (value, index) {
                    if (event.target.value === value.name + " -Artist") {
                        window.location.href = `http://localhost:8080/artistinfo?artist=${value.id}`;
                    }
                });
            });
        })
        .catch(error => console.error(error))
})