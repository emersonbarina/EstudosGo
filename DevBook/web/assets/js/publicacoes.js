$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);

function criarPublicacao(evento) {
  evento.preventDefault();

  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val(),
    }
  }).done(function() {
    window.location = "/home";
  }).fail(function() {
    alert("Erro ao criar a publicação");
  });
}

function curtirPublicacao(evento) {
  // console.log("curtindo publicação...");

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
  // console.log(publicacaoId);

  elementoClicado.prop('disabled', true);

  $.ajax({
      url: `/publicacoes/${publicacaoId}/curtir`,
      method: "POST"
  }).done(function() {
    // alert("Publicação curtida");
    const contadorDeCurtidas = elementoClicado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

  }).fail(function() {
    alert("Erro ao curtir publicação");
  }).always(function() {
    elementoClicado.prop('disabled', false);
  });
}