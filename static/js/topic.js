const add = document.querySelector('.modal-add');
const modal = document.querySelector('.modal-com');
const overlay = document.querySelector('.overlay');
form = document.getElementById('com-form');

function addCom() {
    modal.style.display = 'block';
    overlay.style.display = 'block';
    add.style.display = 'none';
}

function closeCom() {
    modal.style.display = 'none';
    overlay.style.display = 'none';
    add.style.display = 'flex';
}

function submitForm() {
    form.submit();
}