import axios from 'axios';

export const refreshAccessToken = async () => {
    const response = await axios.post(
        "http://localhost:8080/refreshToken",
        {},
        {withCredentials: true}
    )
};
