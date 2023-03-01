const rapButton = document.getElementById("rap-button");
const popvibesButton = document.getElementById("vibes-button");
const electrobutton = document.getElementById("electro-button");
const rockbutton = document.getElementById("rock-button");
const metalbutton = document.getElementById("metal-button");
const hiphopbutton = document.getElementById("hip-hop-button");
const alternativerockbutton= document.getElementById("alternative-rock-button");

rapButton.addEventListener('click', function() {
    let urlParameter = new URLSearchParams();
    urlParameter.append('genre', 'rap');

    let urlNewPage = `http://localhost:8080/categorie?${urlParameter.toString()}`
    window.location.href = urlNewPage
});

popvibesButton.addEventListener('click', function() {
    let urlParameterPop = new URLSearchParams();
    urlParameterPop.append('genre', 'pop-vibes');

    let urlNewPagepop = `http://localhost:8080/categorie?${urlParameterpop.toString()}`
    window.location.href = urlNewPagepop
});

electrobutton.addEventListener('click', function() {
    let urlParameterelectro = new URLSearchParams();
    urlParameterelectro.append('genre', 'electro');

    let urlNewPageelectro = `http://localhost:8080/categorie?${urlParameterelectro.toString()}`
    window.location.href = urlNewPageelectro
});
rockbutton.addEventListener('click', function() {
    let urlParameterrock = new URLSearchParams();
    urlParameterrock.append('genre', 'rock');

    let urlNewPagerock = `http://localhost:8008/categorie?${urlParameterrock.toString()}`
    window.location.href = urlNewPagerock
});
metalbutton.addEventListener('click', function() {
    let urlParametermetal = new URLSearchParams();
    urlParametermetal.append('genre', 'metal');

    let urlNewPagemetal = `http://localhost:8080/categorie?${urlParametermetal.toString()}`
    window.location.href = urlNewPagemetal
});
hiphopbutton.addEventListener('click', function() {
    let urlParameterhiphop = new URLSearchParams();
    urlParameterhiphop.append('genre', 'hip-hop');

    let urlNewPagehiphop = `http://localhost:8080/categorie?${urlParameterhiphop.toString()}`
    window.location.href = urlNewPagehiphop
});
alternativerockbutton.addEventListener('click', function() {
    let urlParameteralternativerock = new URLSearchParams();
    urlParameteralternativerock.append('genre', 'alternative-rock');

    let urlNewPagealternativerock = `http://localhost:8080/categorie?${urlParameteralternativerock.toString()}`
    window.location.href = urlNewPagealternativerock
});