const rapButton = document.getElementById("rap-button");
const popVibesButton = document.getElementById("vibes-button");


rapButton.addEventListener('click', function() {
    let urlParameter = new URLSearchParams();
    urlParameter.append('genre', 'rap');

    let urlNewPage = `http://localhost:3001/categorie?${urlParameter.toString()}`
    window.location.href = urlNewPage
});

popVibesButton.addEventListener('click', function() {
    let urlParameterPop = new URLSearchParams();
    urlParameterPop.append('genre', 'Pop-vibes');

    let urlNewPagePop = `http://localhost:3001/categorie?${urlParameterPop.toString()}`
    window.location.href = urlNewPagePop
});