function autoColorScheme() {
    if (window.matchMedia('(prefers-color-scheme: light)').matches) {
        document.body.classList.add('light');
    } else {
        document.body.classList.add('dark');
    }
}

function setColorScheme(scheme) {
    document.body.classList.add(scheme);
}
