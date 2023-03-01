const rapButton = document.getElementById("rap-button");
const popVibesButton = document.getElementById("vibes-button");
const electrobutton = document.getElementById("electro-button");
const rockbutton = document.getElementById("rock-button");
const metalbutton = document.getElementById("metal-button");

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

electrobutton.addEventListener('click', function() {
    let urlParameterElectro = new URLSearchParams();
    urlParameterElectro.append('genre', 'Electro');

    let urlNewPageElectro = `http://localhost:3001/categorie?${urlParameterElectro.toString()}`
    window.location.href = urlNewPageElectro
});
rockbutton.addEventListener('click', function() {
    let urlParameterRock = new URLSearchParams();
    urlParameterRock.append('genre', 'Rock');

    let urlNewPageRock = `http://localhost:3001/categorie?${urlParameterRock.toString()}`
    window.location.href = urlNewPageRock
});
metalbutton.addEventListener('click', function() {
    let urlParameterMetal = new URLSearchParams();
    urlParameterMetal.append('genre', 'Metal');

    let urlNewPageMetal = `http://localhost:3001/categorie?${urlParameterMetal.toString()}`
    window.location.href = urlNewPageMetal
});