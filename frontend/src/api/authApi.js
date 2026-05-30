import axios from 'axios';
import { API } from './api';

export const refreshAccessToken = async () => {
    const response = await axios.post(
        `${API}/refreshToken`,
        {},
        {withCredentials: true}
    )
};