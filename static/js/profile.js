//----------------------------------------------//
// Supression de topic -------------------------//
//----------------------------------------------//

function deleteTopic(id) {
    var result = confirm("Voulez-vous vraiment supprimer ce topic ?");
    if (result) {
        fetch("/delete-topic", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ PostId: id }),
        })
        .then(response => {
            if (response.ok) {
                // Si la suppression est réussie, masquez le post
                const topicElement = document.getElementById(id);
                topicElement.style.display = "none";
                
                // Décrémentez le nombre de posts affiché
                const nbTopicsElement = document.getElementById("nb-posts");
                const currentNbTopics = parseInt(nbTopicsElement.textContent);
                nbTopicsElement.textContent = currentNbTopics - 1;

                console.log("Topic deleted successfully");
            } else {
                console.error("Failed to delete topic");
            }
        })
        .catch(error => {
            console.error("Error deleting topic:", error);
        });
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

//----------------------------------------------//
// Gestion du formulaire -----------------------//
//----------------------------------------------//

// Récupération de l'image de profil pour pouvoir la mettre à jour
var profilePic = document.getElementById("ProfileImage");

// Permettre à l'utilisateur de sélectionner une nouvelle image et l'enregistrer
var newProfilePic = null;

profilePic.addEventListener("click", function() {
    var input = document.createElement("input");
    input.type = "file";
    input.accept = "image/*";
    input.addEventListener("change", function() {
        newProfilePic = input.files[0];
        var reader = new FileReader();
        reader.onload = function(e) {
            profilePic.src = e.target.result;
        };
        reader.readAsDataURL(newProfilePic);
    });
    input.click();
});

// Envoi du formulaire

function submitForm() {
    username = document.getElementById("username").value;
    email = document.getElementById("email").value;
    desc = document.getElementById("desc").value;
    image = newProfilePic;

    console.log("Username:", username);
    console.log("Email:", email);
    console.log("Description:", desc);
    console.log("Image:", image);

    var formData = new FormData();
    formData.append("username", username);
    formData.append("email", email);
    formData.append("desc", desc);
    formData.append("image", image);

    fetch("/update-profile", {
        method: "POST",
        body: formData,
    })
    .then(response => {
        if (response.ok) {
            console.log("Profile updated successfully"); 
            closeModal();
            location.reload();
        } else {
            console.error("Failed to update profile");
        }
    })
    .catch(error => {
        console.error("Error updating profile:", error);
    });
}

//----------------------------------------------//
// Deconnexion utilisateur ---------------------//
//----------------------------------------------//

function logout() {
    var result = confirm("Voulez-vous vraiment vous déconnecter ?");
    if (result) {
        fetch("/logout", {
            method: "POST",
        })
        .then(response => {
            if (response.ok) {
                console.log("User logged out successfully");
                location.reload();
            } else {
                console.error("Failed to log out user");
            }
        })
        .catch(error => {
            console.error("Error logging out user:", error);
        });
    } else {
        alert("Sérieusement ? J'espère que vous n'êtes pas pareil au lit...")
    }
}