$('#nova-publicacao').on('submit', criarPublicacao);
// $('.curtir-publicacao').on('click', curtirPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#salvar-publicacao').on('click', salvarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

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

    elementoClicado.addClass('descurtir-publicacao');
    elementoClicado.addClass('text-danger');
    elementoClicado.removeClass('curtir-publicacao');

  }).fail(function() {
    alert("Erro ao curtir publicação");
  }).always(function() {
    elementoClicado.prop('disabled', false);
  });
}

function descurtirPublicacao(evento) {

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

  elementoClicado.prop('disabled', true);

  $.ajax({
      url: `/publicacoes/${publicacaoId}/descurtir`,
      method: "POST"
  }).done(function() {
    // alert("Publicação curtida");
    const contadorDeCurtidas = elementoClicado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

    elementoClicado.removeClass('descurtir-publicacao');
    elementoClicado.removeClass('text-danger');
    elementoClicado.addClass('curtir-publicacao');

  }).fail(function() {
    alert("Erro ao curtir publicação");
  }).always(function() {
    elementoClicado.prop('disabled', false);
  });
}

function salvarPublicacao() {
  $(this).prop('disabled', true);

  const publicacaoId = $(this).data('publicacao-id');
  // console.log(publicacaoId);

  $.ajax({
    url: `/publicacoes/${publicacaoId}`,
    method: "PUT",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val()
    }
  }).done(function() {
    alert("Publicação atualizada");
  }).fail(function() {
    alert("Erro ao salvar a publicação");
  }).always(function() {
    $('#salvar-publicacao').prop('disabled', false);
  });
}