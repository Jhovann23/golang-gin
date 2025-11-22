// import useMutation dari '@tanstack/react-query';
import { useMutation } from '@tanstack/react-query';

//import API
import Api from '../../services/api';

//interface RegisterRequest
interface RegisterRequest {
    name: string,
    username: string,
    email: string,
    password: string
}

export const useRegister = () => {
    return useMutation({
        //mutation register
        mutationFn: async (data: RegisterRequest) => {
            //servie API untuk register
            const response = await Api.post('api/register', data);
            return response.data;
        }

    })
}