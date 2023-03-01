const rapButton = document.getElementById("rap-button");

rapButton.addEventListener('click', function() {
    let urlParameter = new URLSearchParams();

    urlParameter.append('genre', 'rap');

    let urlNewPage = `http://localhost:8080/categorie?${urlParameter.toString()}`
    window.location.href = urlNewPage
});