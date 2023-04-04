<script lang="ts">
    import type {User} from "./User";

    let rita : User = {name: "Rita", surname: "Sidorskaya", login: "Klemencya", email:"email", password: "1111", id: 1, preferences: "Italian food"};
    let tanya : User = {name: "Tanya", surname: "Nechepurenko", login: "Tanechka", email:"email", password: "1111", id: 2, preferences: "Russian food"};
    let listOfUsers = [rita, tanya]
</script>

<div class="form">
    <p>Information for users</p>
    <div>
        <p>List of users</p>
        {#each listOfUsers as user}
            <div class="user-info">
                <div>{user.login}</div>
                <div>{user.preferences}</div>
                <div style="float: right"><button>Let's Go Eat</button></div>
            </div>
        {/each}
    </div>
</div>

<style>
    .user-info {
        padding: 10px;
        border-bottom: #2A7D53 4px solid;
        text-align: left;
        display: flex;
        flex-direction: row;
    }

    .user-info div {
        padding: 10px;
    }

    .form {
        text-align: center;
        margin-right: 30%;
        margin-left: 30%;
        margin-top: 10%;
        border: #9e4eca;
    }
</style>
