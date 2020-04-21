export default async function ({store, redirect}) {
  // TODO ナビゲーションガード
  if (!!store.getters.hasLoggedIn) {
    return redirect('/')
  }
}
