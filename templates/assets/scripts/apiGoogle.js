var carte = new google.maps.Map(document.getElementById('map'), {
    center: {lat: 48.8566, lng: 2.3522},
    zoom: 12
});

var marqueur1 = new google.maps.Marker({
    position: {lat: 48.8566, lng: 2.3522},
    map: carte,
    title: 'Paris'
});

var marqueur2 = new google.maps.Marker({
    position: {lat: 51.5074, lng: -0.1278},
    map: carte,
    title: 'Londres'
});

var contenu1 = '<h2>Paris</h2><p>La ville de l\'amour!</p>';
var infobulle1 = new google.maps.InfoWindow({
    content: contenu1
});
marqueur1.addListener('click', function() {
    infobulle1.open(carte, marqueur1);
});

var contenu2 = '<h2>Londres</h2><p>La ville de la royauté!</p>';
var infobulle2 = new google.maps.InfoWindow({
    content: contenu2
});
marqueur2.addListener('click', function() {
    infobulle2.open(carte, marqueur2);
});