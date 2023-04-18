<script lang="ts">
    import type {User} from "./User";
    import type {Request} from "./Request";

    let visibility = false;
    let receiverUser = '';
    let currentUser : User;

    let rita : User = {name: "Rita", surname: "Sidorskaya", login: "Klemencya", email:"email", password: "1111", id: 1, preferences: "Italian food"};
    let tanya : User = {name: "Tanya", surname: "Nechepurenko", login: "Tanechka", email:"email", password: "1111", id: 2, preferences: "Russian food"};
    let listOfUsers = [rita, tanya]

    let request1 : Request = {fromUser: "Klemencya", toUser: "Tanechka", message: "Let's eat some pizza", accept: false};
    let request2 : Request = {fromUser: "Klemencya", toUser: "Tanechka", message: "Let's eat some pizza", accept: false};
    let request3 : Request = {fromUser: "Klemencya", toUser: "Tanechka", message: "Let's eat some pizza", accept: false};
    let listOfRequests = [request1, request2, request3]

    let place = '';
    let cuisine = '';
    let comments = '';

    function openMessageWindow(name: string){
        if (!visibility){
            visibility = !visibility;
            receiverUser = name
        } else {
            receiverUser = name;
        }
    }

    async function sendRequest(){
        let API_URL = 'http://localhost:8080/api/send/request'

        let response = await fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                receiverUser,
                place,
                cuisine,
                comments
            })
        });
        const json = await response.json()
        console.log(json)
    }

    async function getRequestsForUser(){
        let API_URL = 'http://localhost:8080/api/get/request'
        const params = new URLSearchParams({
            login: currentUser.login
        }).toString();

        let response = await fetch(API_URL + params);
        const requests = await response.json();

        return requests
    }

</script>

<div class="form">
    <div id="info-block">
        <h1>Let's Go Eat</h1>
    </div>
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
        </div>
        {#if visibility}
            <div class="element">
                <p>Your message for {receiverUser}</p>
                <div class="message-form">
                    <p><b>Suggest place:</b> <input type="text" bind:value={place} /></p>
                    <p><b>Cuisine:</b> <input type="text" bind:value={cuisine} /></p>
                    <p><b>Comments:</b> <textarea bind:value={comments}></textarea></p>
                </div>

                <button on:click={()=>sendRequest()}>Send request</button>
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
                    </div>
                    <div style="float: right; padding-top: 20px">
                        <button class="accept-button">Let's Go Eat</button>
                    </div>
                </div>
            {/each}
        </div>
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
