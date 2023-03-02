const rapButton = document.getElementById("rap-button");
const popvibesButton = document.getElementById("vibes-button");
const electrobutton = document.getElementById("electro-button");
const rockbutton = document.getElementById("rock-button");
const metalbutton = document.getElementById("metal-button");
const hiphopbutton = document.getElementById("hip-hop-button");
const alternativebutton= document.getElementById("alternative-button");
const reggaebutton= document.getElementById("reggae-button");
const jazzbutton= document.getElementById("jazz-button");

rapButton.addEventListener('click', function() {
    let urlParameter = new URLSearchParams();
    urlParameter.append('genre', 'rap');

    let urlNewPage = `http://localhost:8080/categorie?${urlParameter.toString()}`
    window.location.href = urlNewPage
});

popvibesButton.addEventListener('click', function() {
    let urlParameterPop = new URLSearchParams();
    urlParameterPop.append('genre', 'pop-vibes');

    let urlNewPagePop = `http://localhost:8080/categorie?${urlParameterPop.toString()}`
    window.location.href = urlNewPagePop
});

electrobutton.addEventListener('click', function() {
    let urlParameterElectro = new URLSearchParams();
    urlParameterElectro.append('genre', 'electro');

    let urlNewPageElectro = `http://localhost:8080/categorie?${urlParameterElectro.toString()}`
    window.location.href = urlNewPageElectro
});
rockbutton.addEventListener('click', function() {
    let urlParameterRock = new URLSearchParams();
    urlParameterRock.append('genre', 'rock');

    let urlNewPageRock = `http://localhost:8008/categorie?${urlParameterRock.toString()}`
    window.location.href = urlNewPageRock
});
metalbutton.addEventListener('click', function() {
    let urlParameterMetal = new URLSearchParams();
    urlParameterMetal.append('genre', 'metal');

    let urlNewPageMetal = `http://localhost:8080/categorie?${urlParameterMetal.toString()}`
    window.location.href = urlNewPageMetal
});
hiphopbutton.addEventListener('click', function() {
    let urlParameterHiphop = new URLSearchParams();
    urlParameterHiphop.append('genre', 'hip-hop');

    let urlNewPageHiphop = `http://localhost:8080/categorie?${urlParameterHiphop.toString()}`
    window.location.href = urlNewPageHiphop
});
alternativebutton.addEventListener('click', function() {
    let urlParameterAlternative = new URLSearchParams();
    urlParameterAlternative.append('genre', 'alternative');

    let urlNewPageAlternative = `http://localhost:8080/categorie?${urlParameterAlternative.toString()}`
    window.location.href = urlNewPageAlternative
});
reggaebutton.addEventListener('click', function() {
    let urlParameterReggae = new URLSearchParams();
    urlParameterReggae.append('genre', 'reggae');

    let urlNewPageReggae = `http://localhost:8080/categorie?${urlParameterReggae.toString()}`
    window.location.href = urlNewPageReggae
});
jazzbutton.addEventListener('click', function() {
    let urlParameterJazz = new URLSearchParams();
    urlParameterJazz.append('genre', 'jazz');

    let urlNewPageJazz = `http://localhost:8080/categorie?${urlParameterJazz.toString()}`
    window.location.href = urlNewPageJazz
});

function search() {
    var query = document.getElementById("search-input").value;
    var resultsDiv = document.getElementById("search-results");
    resultsDiv.innerHTML = "Search results for\n : " + query;
    return false;
}
function dropdown() {
    document.getElementById("myDropdown").classList.toggle("show");
}
