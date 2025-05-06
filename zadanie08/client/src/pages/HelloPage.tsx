export default function HelloPage() {
    const handleLogout = () => {
        // wyczyść token / sesję po stronie klienta
        // przekieruj do logowania lub zawołaj endpoint wylogowania
        console.log('Wylogowano');
        window.location.href = '/'; // lub np. navigate('/login') z react-router
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-dark-900">
            <div className="bg-gray-700 p-8 rounded-2xl shadow-md w-full max-w-md text-center">
                <h1 className="text-3xl font-bold mb-4">Witaj w aplikacji!</h1>
                <p className="text-white mb-3">Zostałeś pomyślnie zalogowany.</p>
                <p className="text-stone-50"> Twoje dane to: </p>
                <button
                    onClick={handleLogout}
                    className="mt-10 px-6 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors cursor-pointer"
                >
                    Wyloguj się
                </button>
            </div>
        </div>
    );
}
