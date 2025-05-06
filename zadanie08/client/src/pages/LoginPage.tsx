import { useState } from 'react';
import {Link} from "react-router-dom";

export default function LoginPage() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        // tutaj wyślij dane logowania do backendu
        console.log({ email, password });
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-dark-900">
            <div className="bg-gray-700 p-8 rounded-2xl shadow-md w-full max-w-md">
                <h2 className="text-2xl font-bold mb-6 text-center">Zaloguj się</h2>
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
                    <button
                        type="submit"
                        className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition-colors cursor-pointer"
                    >
                        Zaloguj się
                    </button>
                </form>

                <div className="mt-6 text-center">
                    <Link to="/register">
                        <p className="text-md cursor-pointer underline">Nie masz jeszcze konta? Zarejestruj się!</p>
                    </Link>
                    <p className="text-md mt-2">lub zaloguj się przez:</p>
                    <div className="flex justify-center gap-4 mt-3">
                        <button className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 cursor-pointer">
                            Google
                        </button>
                        <button className="bg-gray-800 text-white px-4 py-2 rounded-lg hover:bg-gray-900 cursor-pointer">
                            GitHub
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}
