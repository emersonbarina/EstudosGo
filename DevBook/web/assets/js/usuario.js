$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);

function pararDeSeguir() {
  const usuarioID = $(this).data('usuario-id');
  $(this).prop('disable', true);

  $.ajax({
    url: `/usuarios/${usuarioID}/parar-de-seguir`,
    method: "POST"
  }).done(function() {
    window.location = `/usuario/${usuarioID}`;
  }).fail(function() {
    FileSystemWritableFileStream.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
    $(`#parar-de-seguir`).prop('disabled', false);
  });
}

function seguir() {

}