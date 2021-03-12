<template>
    <div class="layout">
        <Layout>
            <Header>
                <Menu mode="horizontal" theme="dark" active-name="1"	>
                    <MenuItem name="trace" :to="{ path: '/trace'}">
                        <Icon type="md-analytics" />
                        链路监控
                    </MenuItem>
                    <MenuItem name="2" :to="{path: '/connection' }">
                        <Icon type="md-git-network" />
                        链接查看
                    </MenuItem>
                </Menu>
            </Header>
            <Content>
                <router-view />
            </Content>

        </Layout>
    </div>
</template>

<script>
</script>

<style scoped>
</style>