function submitForm(form) {
    form.submit();
}

function addSubmitLinkListeners() {
    document.querySelectorAll('.submitLink').forEach(link => {
        link.addEventListener('click', function(event) {
            event.preventDefault();
            const form = this.closest('form');
            submitForm(form);
        });
    });
}

// Call the function to add the event listeners
document.addEventListener('DOMContentLoaded', addSubmitLinkListeners);