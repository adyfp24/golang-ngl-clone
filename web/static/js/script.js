document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('send-other').addEventListener('click', function () {
        window.location.href = "/question";
    });
});

document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('submit-message').addEventListener('click', function () {
        document.getElementById('form-message').submit();
    });
});

function redirectToTerms() {
    window.location.href = '/terms';
}

function redirectToSoon() {
    window.location.href = '/soon';
}

function redirectToPrivacy() {
    window.location.href = '/privacy';
}

document.getElementById("download-btn").addEventListener("click", function () {
    html2canvas(document.getElementById("downloaded-card"), {
        useCORS: true
    }).then(function (canvas) {
        var link = document.createElement('a');
        link.download = 'downloaded-card.jpg';
        link.href = canvas.toDataURL("image/jpg");
        link.click();
    });
});

document.getElementById("share-button").addEventListener("click", function () {
    var id = "{{ .question.ID }}"; 
    window.location.href = '/answer/' + id;
});

function validateForm() {
    var question = document.getElementById('question').value;
    var regex = /^[a-zA-Z0-9 .,!?\-]*$/;

    if (!regex.test(question)) {
        alert("Invalid characters in question. Please use only letters, numbers, and punctuation.");
        return false;
    }

    if (question.length > 100) {
        alert("Question cannot exceed 100 characters.");
        return false;
    }

    if (question.trim() === "") {
        alert("Question cannot be empty.");
        return false;
    }

    return true;
}