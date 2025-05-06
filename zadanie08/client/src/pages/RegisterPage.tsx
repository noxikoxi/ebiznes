import {useState} from "react";
import {Link} from "react-router-dom";

const RegisterPage = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [password1, setPassword1] = useState('');
    const [showPasswordError, setShowPasswordError] = useState(false);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (password !== password1) {
            setShowPasswordError(true);
            return;
        }
        setShowPasswordError(false);
        // tutaj wyślij dane logowania do backendu
        console.log({ email, password, password1 });
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-dark-900">
            <div className="bg-gray-700 p-8 rounded-2xl shadow-md w-full max-w-md">
                <h2 className="text-2xl font-bold mb-6 text-center">Rejestracja</h2>
                {showPasswordError && (
                    <p className="text-md text-red-400 font-bold">Hasła muszą być takie same.</p>
                )}
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label htmlFor="email" className="block mb-1 font-medium">Email</label>
                        <input
                            type="email"
                            id="email"
                            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            required
                        />
                    </div>
                    <div>
                        <label htmlFor="password" className="block mb-1 font-medium">Hasło</label>
                        <input
                            type="password"
                            id="password"
                            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required
                        />
                    </div>
                    <div>
                        <label htmlFor="password1" className="block mb-1 font-medium">Powtórz hasło</label>
                        <input
                            type="password"
                            id="password1"
                            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                            value={password1}
                            onChange={(e) => setPassword1(e.target.value)}
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition-colors cursor-pointer"
                    >
                        Zarejestruj się
                    </button>
                </form>
                <div className="mt-6 text-center">
                    <Link to="/">
                        <p className="text-md">Masz już konto lub wolisz zalogować się przez Google/Github?</p>
                        <p className="text-md cursor-pointer underline">Zaloguj się!</p>
                    </Link>
                </div>
            </div>
        </div>
    )
}

export default RegisterPage;