document.addEventListener('DOMContentLoaded', function () {
    fetch('/post/searchbar')
        .then(response => response.json())
        .then(data => {
            let valueSearch = document.getElementById("search-input")
            let button = document.querySelector(".search-button");
            valueSearch.addEventListener("input", function() {
                data.artists.forEach(function (value, index) {
                    if (valueSearch.value === value.name) {
                        window.location.href = "http://localhost:8080/artistinfo?artist=damso"
                    }
                });
            })
        })
        .catch(error => console.error(error))
})