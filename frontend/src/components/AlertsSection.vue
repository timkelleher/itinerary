<template>
    <div>
        <div class="row" v-if="this.errorMessage">
            <div class="col-sm-12">
                <div class="alert alert-danger" role="alert">
                    {{ this.errorMessage }}
                </div>
            </div>
        </div>
        <div class="row" v-if="this.successMessage">
            <div class="col-sm-12">
                <div class="alert alert-success" role="alert">
                    {{ this.successMessage }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "AlertsSection",

    props: {
        errorMessage: String,
        successMessage: String,
    }
}
</script>