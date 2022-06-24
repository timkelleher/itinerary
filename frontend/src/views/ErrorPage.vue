<template>
    <div class="col-md-10 offset-md-1 py-2">
        <div class="row">
            <div class="col-8">
                <h3>Error!</h3>
            </div>
            <div class="col-4 text-right">
                <router-link :to="{ name: 'collection_list' }">
                    <button class="btn btn-secondary btn-sm">Back</button>
                </router-link>
            </div>
        </div>
        <div>
            <div class="alert alert-danger" role="alert">
                {{ this.errorMessage }}
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ErrorPage',

    props: [
        'errorMessage'
    ],
}
</script>