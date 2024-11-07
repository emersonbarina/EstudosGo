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
    Swal.fire("Ops...", "Erro ao criar a publicação!", "error");
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
    const contadorDeCurtidas = elementoClicado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

    elementoClicado.addClass('descurtir-publicacao');
    elementoClicado.addClass('text-danger');
    elementoClicado.removeClass('curtir-publicacao');

  }).fail(function() {
    Swal.fire("Ops...", "Erro ao curtir publicação!", "error");
  }).always(function() {
    elementoClicado.prop('disabled', false);
  });
}

function descurtirPublicacao(evento) {
  evento.preventDefault();

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

  elementoClicado.prop('disabled', true);

  $.ajax({
      url: `/publicacoes/${publicacaoId}/descurtir`,
      method: "POST"
  }).done(function() {
    const contadorDeCurtidas = elementoClicado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

    elementoClicado.removeClass('descurtir-publicacao');
    elementoClicado.removeClass('text-danger');
    elementoClicado.addClass('curtir-publicacao');

  }).fail(function() {
    Swal.fire("Ops...", "Erro ao descurtir publicação!", "error");
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
    Swal.fire(
      "Sucesso!", "Publicação atualizada!", "success").then(function() {
        window.location = "/home";
      });
  }).fail(function() {
    Swal.fire("Ops...", "Erro ao salvar a publicação!", "error");
  }).always(function() {
    $('#salvar-publicacao').prop('disabled', false);
  });
}

function deletarPublicacao(evento) {
  evento.preventDefault();

  Swal.fire({
    title: "Atenção!",
    text: "Ação irreversível, deseja realmente excluir essa publicação?",
    showCancelButton: true,
    cancelButtomText: "Cancelar",
    icon: "warning"
  }).then(function(confirmacao) {
    if (!confirmacao.value) return;

    const elementoClicado = $(evento.target);
    const publicacao = elementoClicado.closest('div');
    const publicacaoId = publicacao.data('publicacao-id');
  
    elementoClicado.prop('disabled', true);
  
    $.ajax({
      url: `/publicacoes/${publicacaoId}`,
      method: "DELETE"
    }).done(function() {
      publicacao.fadeOut("slow", function() {
        $(this).remove();
      });
    }).fail(function() {
      Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
    }).always(function() {
      elementoClicado.prop('disabled', false);   
    })
  })
}