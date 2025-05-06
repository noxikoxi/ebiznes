import { useEffect, useState } from "react";
import axios from "axios";

export interface UserData {
    email : string,
    name: string,
    surname: string
}

export const useAuth = () => {
    const [user, setUser] = useState<UserData>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios.get("http://localhost:1323/user", { withCredentials: true })
            .then(res => {
                setUser(res.data)
            })
            .catch(() => setUser(null))
            .finally(() => setLoading(false));
    }, []);

    return { user, loading };
};
