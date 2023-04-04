<script lang="ts">
    import type {User} from "./User";
    import {field, form} from "svelte-forms";
    import {required} from "svelte-forms/validators";

    let visibility = false;
    let recieverName = ''

    let rita : User = {name: "Rita", surname: "Sidorskaya", login: "Klemencya", email:"email", password: "1111", id: 1, preferences: "Italian food"};
    let tanya : User = {name: "Tanya", surname: "Nechepurenko", login: "Tanechka", email:"email", password: "1111", id: 2, preferences: "Russian food"};
    let listOfUsers = [rita, tanya]

    const place = field('place', '', [required()]);
    const cuisine = field('cuisine', '', [required()]);
    const comments = field('comments', '', [required()]);
    const myForm = form(place, cuisine, comments);

    function openMessageWindow(name: string){
        if (!visibility){
            visibility = !visibility;
            recieverName = name
        } else {
            recieverName = name;
        }
    }
</script>

<div class="form">
    <div id="info-block">
        <h1>Let's Go Eat</h1>
        <p>This app has been specially designed for people who can't find company to eat.</p>
        <p>Just choose your company and eat</p>
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
                <p>Your message for {recieverName}</p>
                <div class="message-form">
                    <p><b>Suggest place:</b> <input type="text" bind:value={$place.value} /></p>
                    <p><b>Cuisine:</b> <input type="text" bind:value={$cuisine.value} /></p>
                    <p><b>Comments:</b> <textarea bind:value={$comments.value}></textarea></p>
                </div>

                <button disabled={!$myForm.valid}>Send request</button>
            </div>
        {/if}
    </div>
</div>

<style>
    .user-info {
        padding: 10px;
        border-bottom: #2A7D53 4px solid;
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
