function autoColorScheme() {
    if (window.matchMedia('(prefers-color-scheme: light)').matches) {
        document.body.classList.add('light');
    } else {
        document.body.classList.add('dark');
    }
}

function toggleColorScheme(scheme) {
    document.body.classList.add(scheme);
}

window.onload = autoColorScheme()
