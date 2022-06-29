import axios from "axios";
import { Auth } from "aws-amplify";

const domain = "https://api.pomo.cafe"

async function getToken() {
    let session = await Auth.currentSession();
    let token = session.getIdToken().getJwtToken();

    return token
}

async function getUsername() {
        let authUser = await Auth.currentAuthenticatedUser();
        let username = authUser.getUsername();

        return username
}

export async function userAPI(method, path, body) {
    let resp = await axios({
        method,
        url: `${domain}/${await getUsername() + path}`,
        headers: { Authorization: "Bearer " + await getToken() },
        data: body,
    });

    // TODO remove
    console.log(await getUsername());

    return resp.data;
}