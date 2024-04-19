//----------------------------------------------//
// Barre de recherche --------------------------//
//----------------------------------------------//

function HideAndShow() {
    var searchValue = document.getElementById("searchbar").value.toLowerCase();
    var topics = document.querySelectorAll(".post");
    var noresults = document.getElementById("no-results");
    var visibleTopicsCount = 0; 

    topics.forEach(function(topic) {
        var title = topic.querySelector(".title").textContent.toLowerCase();
        var username = topic.querySelector(".username").textContent.toLowerCase();

        if (title.includes(searchValue) || username.includes(searchValue)) {
            topic.style.display = "flex";
            visibleTopicsCount++; 
        } else {
            topic.style.display = "none";
        }
    });

    // Vérifier si aucun sujet n'est visible
    if (visibleTopicsCount === 0) {
        noresults.style.display = "block";
    } else {
        noresults.style.display = "none"; 
    }
}

//----------------------------------------------//
// Gestion du modal ----------------------------//
//----------------------------------------------//

// Récupération du modal et de l'overlay
var modal = document.getElementById("modal");
var overlay = document.getElementById("overlay");

// Fonction pour ouvrir et fermer le modal
function openModal() {
    modal.style.display = "block";
    overlay.style.display = "block";
}

function closeModal() {
    modal.style.display = "none";
    overlay.style.display = "none";
}

// Fermer le modal si l'utilisateur clique en dehors de celui-ci ou sur le bouton de fermeture
window.onclick = function(event) {
    if (event.target == overlay) {
        closeModal();
    }
}

// Envoyer le formulaire
document.getElementById("submitBtn").addEventListener("click", function() {
    document.getElementById("myForm").submit();
});

//----------------------------------------------//
// Gestion de l'ouverture des topics -----------//
//----------------------------------------------//

function openTopic(id) {
    window.location.href = "/topic/" + id;
}