import {useNavigate} from "react-router-dom";
import { useLocation } from "react-router-dom";

export default function HelloPage() {
    const navigate = useNavigate();
    const location = useLocation();
    const { email, name, surname } = location.state || {};

    const handleLogout = () => {
        navigate("/login")
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-dark-900">
            <div className="bg-gray-700 p-8 rounded-2xl shadow-md w-full max-w-md text-center">
                <h1 className="text-3xl font-bold mb-4 text-xxl">Witaj w aplikacji!</h1>
                <p className="text-white mb-3 text-xl">Zostałeś pomyślnie zalogowany.</p>
                <p className="text-stone-50 text-xl"> Twoje dane to: </p>
                <div className="flex flex-col">
                    <span className="text-xl"><span className="text-stone-50 font-bold ">Email:</span> {email}</span>
                    <span className="text-xl"><span className="text-stone-50 font-bold ">Imię:</span> {name}</span>
                    <span className="text-xl"><span className="text-stone-50 font-bold ">Nazwisko:</span> {surname}</span>
                </div>
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
