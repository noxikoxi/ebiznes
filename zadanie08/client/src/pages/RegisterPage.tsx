import {useState} from "react";
import {Link, useNavigate} from "react-router-dom";
import axios from "axios";

const RegisterPage = () => {
    const [email, setEmail] = useState('');
    const [name, setName] = useState('');
    const [surname, setSurname] = useState('');
    const [password, setPassword] = useState('');
    const [password2, setPassword2] = useState('');
    const [showPasswordError, setShowPasswordError] = useState(false);
    const [error, setError] = useState<string | null>(null)
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(null)
        setLoading(true)
        if (password !== password2) {
            setShowPasswordError(true);
            setLoading(false)
            return;
        }
        setShowPasswordError(false);
        const userData = {
            email,
            name,
            surname,
            password
        };
        try {
            await axios.post('http://localhost:1323/register', userData);
            navigate("/login");
        } catch (error: any) {
            if (error.response) {
                setError(error.response.data.error || `HTTP error! status: ${error.response.status}`);
            } else if (error.request) {
                setError('No response from server');
            } else {
                setError(error.message || 'An error occurred');
            }
            console.error('Axios error:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-dark-900">
            <div className="bg-gray-700 p-8 rounded-2xl shadow-md w-full max-w-md">
                <h2 className="text-2xl font-bold mb-6 text-center">Rejestracja</h2>
                {showPasswordError && (
                    <p className="text-md text-red-400 font-bold">Hasła muszą być takie same.</p>
                )}
                {error && <p className="text-md text-red-400 font-bold">Error: {error}</p>}
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
                        <label htmlFor="name" className="block mb-1 font-medium">Imię</label>
                        <input
                            type="text"
                            id="name"
                            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            required
                        />
                    </div>
                    <div>
                        <label htmlFor="surname" className="block mb-1 font-medium">Nazwisko</label>
                        <input
                            type="text"
                            id="surname"
                            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                            value={surname}
                            onChange={(e) => setSurname(e.target.value)}
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
                            value={password2}
                            onChange={(e) => setPassword2(e.target.value)}
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition-colors cursor-pointer"
                        disabled={loading}
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