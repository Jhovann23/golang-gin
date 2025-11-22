// import hook react
import { useContext } from "react";

// import js-cookie
import Cookies from "js-cookie";

// import hook useNavigate dari react-router
import { useNavigate } from "react-router";

// import context
import { AuthContext } from "../../context/AuthContext";

export const useLogout = (): (() => void) => {
    // Ambil setIsAuthenticated dari context
    const authContext = useContext(AuthContext);
    // Gunakan null assertion karena kita yakin AuthContext akan selalu tersedia
    const {setIsAuthenticated} = authContext!;
    //inialisasi navigate
    const navigate = useNavigate();

    //fungsi logout
    const logout = (): void => {
        //hapus token
        Cookies.remove('user');
        Cookies.remove('token')
        //set authenticated jadi false
        setIsAuthenticated(false);
        //navigate ke login
        navigate('/login')
    }

    return logout;
}