<script lang="ts">
    import type {User} from "./User";
    import type {Request} from "./Request";
    import {onMount} from "svelte";
    import {field, form} from "svelte-forms";
    import {required} from "svelte-forms/validators";
    import {page} from '$app/stores';

    let visibility = false;
    let receiverUser = '';
    let currentUser: User;
    let login = getParameterByName('user');

    let listOfUsers : [User] = [];
    let listOfRequests : [Request] = [];

    let place = field('place', '', [required()]);
    let cuisine = field('cuisine', '', [required()]);
    let invitation = field('invitation', '', []);
    let myForm = form(place, cuisine, invitation);

    // This method will separate current User nickname from url
    function getParameterByName(name) {
        let url = $page.url;
        // console.log("url " + url)
        name = name.replace(/[\[\]]/g, '\\$&');
        let regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
            results = regex.exec(url.toString());
        if (!results) return null;
        if (!results[2]) return '';
        return decodeURIComponent(results[2].replace(/\+/g, ' '));
    }

    onMount(async () => {
        let API_URL = 'http://localhost:8080/api/get/user'

        const fetchResponse = fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                login: login
            })
        });
        fetchResponse
            .then(response => response.json())
            .then(data => {
                const userInfo = JSON.parse(data.users)
                currentUser = {
                    name: userInfo.Name,
                    surname: userInfo.Surname,
                    login: userInfo.Login,
                    email: userInfo.Email,
                    password: userInfo.Password,
                    id: userInfo.ID,
                    preferences: userInfo.Preferences
                }

            });

        await getAllUsers();
        await getRequestsForUser();
    })

    async function getAllUsers() {
        let API_URL = 'http://localhost:8080/api/get/allusers'
        let listOfFriends : [User] = []

        const fetchResponse = fetch(API_URL, {method: 'POST'});
        fetchResponse.then(response => response.json())
            .then(data => {
                const usersInfo = JSON.parse(data.users)
                usersInfo.forEach(user => {
                    if (currentUser.login != user.Login) {
                        listOfFriends.push({
                            name: user.Name,
                            surname: user.Surname,
                            login: user.Login,
                            email: user.Email,
                            password: user.Password,
                            id: user.ID,
                            preferences: user.Preferences
                        })
                    }
                })
                listOfUsers = listOfFriends
            })
    }


    function openMessageWindow(name: string) {
        if (!visibility) {
            visibility = !visibility;
            receiverUser = name
        } else {
            receiverUser = name;
        }
    }

    async function sendRequest() {
        let API_URL = 'http://localhost:8080/api/send/request'

        let response = await fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                sender_user: currentUser.login,
                receiver_user: receiverUser,
                place: $place.value,
                cuisine: $cuisine.value,
                invitation: $invitation.value
            })
        });
        await response.json()
    }

    async function getRequestsForUser() {
        let API_URL = 'http://localhost:8080/api/get/request'

        let fetchResponse = fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                login: login
            })
        });
        let listOfFriendRequests : [Request] = []
        fetchResponse.then(response => response.json())
            .then(data => {
                const requestsInfo = JSON.parse(data.users)
                if (requestsInfo != null) {
                    requestsInfo.forEach(request => listOfFriendRequests.push({
                        fromUser: request.Sender,
                        toUser: request.Receiver,
                        message: 'Place: ' + request.Place + '\n Cuisine: ' + request.Cuisine + '\n Comments: ' + request.Invitation,
                        accept: request.Accepted,
                        id: request.ID
                    }))

                }
                listOfRequests = listOfFriendRequests
            })
    }

    async function deleteRequest(requestId: string) {
        let API_URL = 'http://localhost:8080/api/delete/request'

        let fetchResponse = fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                id: requestId
            })
        });

        fetchResponse.then(response => response.json())
            .then(data => {
                if (data.msg == "OK") {
                    getRequestsForUser();
                }
            })
    }

    async function changeRequestAccept(requestId: string) {
        let API_URL = 'http://localhost:8080/api/change/request/status'

        let fetchResponse = fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                id: requestId
            })
        });

        fetchResponse.then(response => response.json())
            .then(data => {
                console.log(data)
                if (data.msg == "OK") {
                    getRequestsForUser();
                }
            })
    }

</script>

<head>
    <title>${login}'s user page</title>
    <meta property="og:title" content="${login}'s user page"/>
    <meta property="og:description" content="Text and invite ${login} to eat together!"/>
</head>

<div class="form">
    <div id="info-block">
        <h1>Let's Go Eat</h1>
    </div>
    <div><p>Hello {login}!</p></div>
    <div class="meeting-form">
        <div class="element">
            <p>List of users</p>
            {#each listOfUsers as user}
                <div class="user-info">
                    <div>
                        <div>{user.login}</div>
                        <div>{user.preferences}</div>
                    </div>
                    <div style="float: right; padding-top: 20px">
                        <button on:click={() => openMessageWindow(user.login)}>Let's Go Eat</button>
                    </div>
                </div>
            {/each}
            <button on:click={()=>getAllUsers()}>Find more</button>
        </div>
        {#if visibility}
            <div class="element">
                <p>Your message for {receiverUser}</p>
                <div class="message-form">
                    <p><b>Suggest place:</b> <input type="text" bind:value={$place.value}/></p>
                    <p><b>Cuisine:</b> <input type="text" bind:value={$cuisine.value}/></p>
                    <p><b>Comments:</b> <textarea bind:value={$invitation.value}></textarea></p>
                </div>
                {#if $myForm.valid}
                    <button on:click={()=>sendRequest()}>Send request</button>
                {/if}
            </div>
        {/if}
    </div>
    <div id="meeting-box">
        <p>Requests</p>
        <div>
            {#each listOfRequests as request}
                <div class="user-info">
                    <div>
                        <div>{request.fromUser}</div>
                        <div>{request.message}</div>
                        <div>{request.accept}</div>
                    </div>
                    <div style="float: right; padding-top: 20px;">
                        {#if !request.accept}
                            <div>
                                <button on:click={()=>changeRequestAccept(request.id)}>Let's Go Eat</button>
                            </div>
                        {:else}
                            <div>
                                <button style="background-color: #85d2ac"
                                        on:click={()=>changeRequestAccept(request.id)}>Let's Go Eat
                                </button>
                            </div>
                        {/if}
                        <div>
                            <button on:click={() => deleteRequest(request.id)}>Delete</button>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
        <button on:click={()=>getRequestsForUser()}>Refresh</button>
    </div>
</div>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Shrikhand&family=Ubuntu:wght@400;700&display=swap');

    h1 {
        font-size: 64px;
        font-family: "Shrikhand", cursive;
        color: #9e4eca;
    }

    #meeting-box {
        overflow: scroll;
        height: 400px;
        margin-top: 10%;
        border: 2px solid #9e4eca;
        padding: 50px
    }

    p {
        font-size: 20px;
        font-family: 'Ubuntu', sans-serif;
    }

    .user-info {
        padding: 10px;
        border-bottom: #9e4eca 4px solid;
        text-align: left;
        display: inline-block;
        width: 100%;
    }

    .message-form {
        text-align: left;
    }

    input {
        width: 100%;
    }

    textarea {
        width: 100%;
    }

    .user-info div {
        display: inline-flex;
    }

    .user-info div div {
        padding-top: 20px;
        padding-right: 50px;
    }

    .form {
        text-align: center;
        margin-right: 10%;
        margin-left: 10%;
        margin-top: 10%;
        border: #9e4eca;
    }

    #info-block {
        padding-bottom: 50px;
    }

    .meeting-form {
        display: flex;
    }

    .element {
        text-align: center;
        margin: 50px 100px;
        width: 100%;
    }

    .element p {
        padding: 10px;
    }
</style>
