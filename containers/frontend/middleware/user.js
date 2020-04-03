export default async function ({store}) {
  console.log("middleware/user before")
  await store.dispatch('auth/currentUser')
  console.log("middleware/user after")
}
