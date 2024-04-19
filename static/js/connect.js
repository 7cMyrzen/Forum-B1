const accountExistElement = document.getElementById("accountExist");
const changeModeButtonElement = document.getElementById("changeModeButton");
const hoverZone = document.querySelector('.hover-zone');

function changeConnectMode() {
    if (hoverZone.style.left === '0%') {
        hoverZone.style.left = '50%';
        accountExistElement.textContent = "Vous n'avez pas encore de compte ?";
        changeModeButtonElement.textContent = "Inscrivez-vous";
    } else {
        hoverZone.style.left = '0%';
        changeModeButtonElement.textContent = "Connectez-vous";
        accountExistElement.textContent = "Vous avez déjà un compte ?";
    }
}