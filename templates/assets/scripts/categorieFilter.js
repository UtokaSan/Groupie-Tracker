let data = [
    { nom: "Groupe 1", membres: 10, dateCreation: "2022-01-01" },
    { nom: "Groupe 2", membres: 5, dateCreation: "2021-05-01" },
    { nom: "Groupe 3", membres: 25, dateCreation: "2021-12-01" },
    { nom: "Groupe 4", membres: 15, dateCreation: "2022-02-01" },
];
function filterData(data, minDate) {
    return data.filter((item) =>

        new Date(item.dateCreation) >= new Date(minDate)
    );
}

document.querySelector("#date-min").addEventListener("change", filterGroups);

function filterGroups() {
    let minDate = document.querySelector("#date-min").value;

    let filteredData = filterData(data, minDate);
    console.log(filteredData);
}