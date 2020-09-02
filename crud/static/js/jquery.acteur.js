/**
	jQuery plugin
	Copyright (C) 2009 Philippe Billerot

 */

(function($) {
	//PLUGIN ROWSELECT d'une vue
	$.fn.rowselect = function(options) 
	{	
		// définition des paramètres par défaut
		var defaults = {
				nomCookie: "rowselect",
				classCss: "rowselect"
		};	
		// mélange des paramètres fournis et des parametres par défaut
		var opts = $.extend(defaults, options);		

		var init = function () {
			// positionnement sur la ligne sélectionnée avant l'actualisation
			//alert('init');
			if ( $.cookie(opts.nomCookie) != null && $.cookie(opts.nomCookie).length > 0 ) {
				var cook = $('#' + $.cookie(opts.nomCookie));
				if ( cook != null ) {
					cook.addClass(opts.classCss);

					var anchor_url = "" + $.cookie(opts.nomCookie);
					Metronic.scrollTo($("#" + anchor_url)
							, -$(".tableFloatingHeaderOriginal").height()-$(".tabbable-line").height());
				} // endif
			} // endif
		}

		function maFonction(obj) {
			obj.click( function(event) {
				//alert('clic');
				if ( obj.attr('class') != opts.classCss ) {
					obj.addClass(opts.classCss);
					$.cookie(opts.nomCookie, obj.attr('id'));
				} // endif
				// suppression sélection autre ligne
				$('.' + opts.classCss).each( function() { 
					if ( $(this).attr('id') != obj.attr('id') ) {
						$(this).removeClass(opts.classCss);
					} // endif 
				});
			}); // end click
			return obj;
		}

		init();

		// boucle sur tous les éléments de l'objet jQuery
		return this.each ( function() {
			maFonction($(this)); 
		}	
		);
	}; // end ROWSELECT

	/*
	 * Traitement d'un champ de recherche dans le portail
	 */
	var menuContent = false;
	$( '.portlet' ).searchable({
		searchField: '#searchable-field',
		selector: 'li',
		childSelector: '.desc',
		show: function( elem ) {
			//console.log('show:' + elem.html());
			var portlet = elem.closest('.portlet');
			
			// Le bloc sera déplié
			portlet.show();
			portlet.find('.portlet-title').find('a').removeClass('expand');
			portlet.find('.portlet-title').find('a').addClass('collapse');
			portlet.find('.portlet-body').show();
			
			// le menu correspondant sera en gras non sélectionné
			var menu = $('#id' + portlet.attr('id'));
			menu.addClass('bold');
			menu.addClass('act-jquery-sidebar-espace');
			menu.closest('li').removeClass('active');
			menu.children('i').removeClass('fa-check-square-o');
			menu.children('i').addClass('fa-square-o');
			
		    // bouton "tout plier / déplier"
			if ( $('#idtoolbar[style="display: none"]') ) {
				$('#idtoolbar').show();
				// les blocs seront dépliés
				$('#idplierdeplier').children('i').removeClass('fa-angle-up');
				$('#idplierdeplier').children('i').addClass('fa-angle-down');
				$('#idplierdeplier').children('span').text('Tout plier');
			}
			
			elem.slideDown(100);
		},
		hide: function( elem ) {
			elem.slideUp( 100 );
		},
		onSearchActive : function( elem, term ) {
			if ( ! menuContent ) {
				menuContent = true;
				$('#idtuile').hide();
				$('#idespace').show();
				$('#id-sidebar-espace').show();
			    $('#searchable-lien').val('');
			    $('#searchable-lien').keyup();
			}
			var portlet = elem.closest('.portlet');
			var menu = $('#id' + portlet.attr('id'));
			menu.removeClass('bold');
			menu.removeClass('act-jquery-sidebar-espace');
			menu.closest('li').removeClass('active');
			menu.children('i').removeClass('fa-check-square-o');
			menu.children('i').addClass('fa-square-o');
			elem.hide();

			if ( $('.act-jquery-sidebar-lien').closest('ul').find('li').hasClass('active'))
		    	$('.act-jquery-sidebar-lien').closest('ul').find('li').removeClass('active');

		},
		onSearchEmpty: function( elem ) {
	        // On réaffiche tout - équivalent à act-espace-raz
			var portlet = elem.closest('.portlet');
			portlet.show();
			var menu = $('#id' + portlet.attr('id'));
			menu.addClass('bold');
			menu.addClass('act-jquery-sidebar-espace');
			menu.closest('li').removeClass('active');
			menu.children('i').removeClass('fa-check-square-o');
			menu.children('i').addClass('fa-square-o');
		
	        elem.show();
	        
	    }
	})
	
	/*
	 * Raz de la recherche et des filtres 
	 */
	$(document).on('click', '#idraz', function (event) {
		
//		$('#idtuile').hide();
//		$('#idespace').show();
//		$('#id-sidebar-espace').show();
		$('#idtuile').show();
		$('#idespace').hide();
		$('#id-sidebar-espace').hide();
		
	    // toolbar
	    $('#idtoolbar').hide();
	    // les blocs seront dépliés
		$('#idplierdeplier').children('i').removeClass('fa-angle-up');
		$('#idplierdeplier').children('i').addClass('fa-angle-down');
		$('#idplierdeplier').children('span').text('Tout plier');

		
		$('.act-espace').show();
		$('.act-jquery-sidebar-lien').closest('li').removeClass('active');
		$('.act-sidebar-espace').addClass('bold');
		$('.act-sidebar-espace').addClass('act-jquery-sidebar-espace');
		$('.act-sidebar-espace').closest('li').removeClass('active');
		$('.act-sidebar-espace').children('i').removeClass('fa-check-square-o');
		$('.act-sidebar-espace').children('i').addClass('fa-square-o');

		$('.act-espace').find('.portlet-title').find('a').removeClass('expand');
		$('.act-espace').find('.portlet-title').find('a').addClass('collapse');
		$('.act-espace').find('.portlet-body').show();

		$('#searchable-field').val('');
	    $('#searchable-field').keyup();
	    
	    $('#searchable-field').focus();
		event.preventDefault();
	})
	
	$( '.act-espace' ).searchable({
		searchField: '#searchable-lien',
		selector: 'li',
		childSelector: '.act-typeapp',
		show: function( elem ) {
			// affichage de l'espace
			var container = elem.closest('.act-espace');
			container.show();
			
			// affichage de l'espace dans le menu
			var portlet = container.find('.portlet');
			var menu = $('#id' + portlet.attr('id'));
			menu.addClass('bold');
			menu.addClass('act-jquery-sidebar-espace');
			
			// Le bloc sera plié
			portlet.show();
			portlet.find('.portlet-title').find('a').removeClass('collapse');
			portlet.find('.portlet-title').find('a').addClass('expand');
			portlet.find('.portlet-body').hide();
			
			elem.slideDown(100);
			
		},
		hide: function( elem ) {
			elem.slideUp( 100 );
		},
		onSearchActive : function( elem, term ) {
		    elem.hide();
			// on enlève le gras des espaces dans le menu
			var portlet = elem.closest('.act-espace').find('.portlet');
			var menu = $('#id' + portlet.attr('id'));
			menu.removeClass('bold');
			menu.removeClass('act-jquery-sidebar-espace');
		    menu.find('i').removeClass('fa-check-square-o');
		    menu.find('i').addClass('fa-square-o');
		},
		onSearchEmpty: function( elem ) {
	        elem.show();
	    }
	})

	/*
	 * Sélection d'une tuile
	 */
	$(document).on('click', '.act-tuile', function (event) {
		$('#idtuile').hide();
		$('#idespace').show();
		$('#id-sidebar-espace').show();
		
		var idespace = $(this).data('espace');
		$('#idtab_'+idespace).click();
	})

	/*
	 * Sélection d'un espace
	 */
	$(document).on('click', '.act-jquery-sidebar-espace', function (event) 
	{
	    var id_target = $(this).data('target');
	    
	    $('#idtuile').hide();
	    $('#idtoolbar').hide();
	    $('#idespace').show();
	    $('#id-sidebar-espace').show();
	    
		if ( $(this).closest('li').hasClass('active') ) {
			$('#idraz').click();
			return;
		}
		// sélection de l'espace
	    $(this).closest('ul').find('li').removeClass('active');
	    $(this).closest('li').addClass('active');
	    
	    // on change l'icone fa-square-o en fa-check-square-o
	    $(this).closest('ul').find('i').removeClass('fa-check-square-o');
	    $(this).closest('ul').find('i').addClass('fa-square-o');
	    $(this).closest('a').find('i').removeClass('fa-square-o');
	    $(this).closest('a').find('i').addClass('fa-check-square-o');

		// on cache tous les blocs
	    $(id_target).closest('.row').find('.portlet').hide();

		// on affiche le bloc sélectionné
	    $(id_target).closest('.portlet').show();
	    
		// expand collapse du bloc
	    $(id_target).find('.portlet-title').find('a').removeClass('expand');
	    $(id_target).find('.portlet-title').find('a').addClass('collapse');
	    $(id_target).find('.portlet-body').show();
	});
	
	/*
	 * Sélection d'un type de lien
	 */
	$(document).on('click', '.act-jquery-sidebar-lien', function (event) 
	{
		if ( $(this).closest('li').hasClass('active') ) {
			$('#idraz').click();
			return;
		}

		$('#idtuile').hide();
	    $('#idespace').show();
	    $('#id-sidebar-espace').show();

	    // toolbar
	    $('#idtoolbar').show();
	    // les blocs seront pliés
		$('#idplierdeplier').children('i').removeClass('fa-angle-down');
		$('#idplierdeplier').children('i').addClass('fa-angle-up');
		$('#idplierdeplier').children('span').text('Tout déplier');

		$(this).closest('ul').find('li').removeClass('active');
	    $(this).closest('li').addClass('active');
	    
	    var espace = $(this).data('espace');
		$('.portlet').removeClass('hide');
		$('.act-jquery-sidebar-espace').addClass('bold');
		$('.act-jquery-sidebar-espace').closest('ul').find('li').removeClass('active');

	    $('#searchable-field').val('');
	    $('#searchable-field').keyup();

	    $('#searchable-lien').val(espace);
	    $('#searchable-lien').keyup();
	    
	});

	/*
	 * Plier / Déplier
	 */
	$(document).on('click', '#idplierdeplier', function (event) {
		if ( $(this).children('i').hasClass('fa-angle-up') ) {
			// les blocs sont pliés
			
			// on change le bouton
			$(this).children('i').removeClass('fa-angle-up');
			$(this).children('i').addClass('fa-angle-down');
			$(this).children('span').text('Tout plier');
			// on déplie les blocs
			$('.act-espace').find('.portlet-title').find('a').removeClass('expand');
			$('.act-espace').find('.portlet-title').find('a').addClass('collapse');
			$('.act-espace').find('.portlet-body').show();

		} else {
			// les blocs sont dépliés
			
			// on change le bouton
			$(this).children('i').removeClass('fa-angle-down');
			$(this).children('i').addClass('fa-angle-up');
			$(this).children('span').text('Tout déplier');
			// on plie les blocs
			$('.act-espace').find('.portlet-title').find('a').removeClass('collapse');
			$('.act-espace').find('.portlet-title').find('a').addClass('expand');
			$('.act-espace').find('.portlet-body').hide();

		}
		
		event.preventDefault();
	})

	/*
	 * Mene des applications dans ACTEUR STUDIO
	 */
	$(document).on('click', '.act-jquery-sidebar-menu', function (event) 
	{
	    var id_target = $(this).data('target');

		if ( $(this).closest('li').hasClass('active') ) {
			// on désélectionne l'espace en cours
			$(this).closest('li').removeClass('active');
			// on affiche tous les blocs
			$(id_target).closest('.row').find('.portlet').show();
			return;
		}

	    $(this).closest('ul').find('li').removeClass('active');
	    $(this).closest('li').addClass('active');

		// on cache tous les blocs
	    $(id_target).closest('.row').find('.portlet').hide();

		// on affiche le bloc sélectionné
	    $(id_target).closest('.portlet').show();

	});

	// MODAL
	var $modal = $('#ajax-modal');
	$('.act-jquery-modal').on('click', function() {
		//url = $(this).data('url');
		$url = 'http://localhost/actigniter/assets/bootstrap-modal/modal_ajax_test.html';
		// create the backdrop and wait for next modal to be triggered
		$('body').modalmanager('loading');

		setTimeout(function() {
			$modal.load($url, '', function() {
				$modal.modal({
					width: '85%',
					height: '85%',
					maxHeight: '85%'
				});
			});
		}, 100);
	});

	// LOGON
	$(document).on('click', '.act-jquery-bootbox', function(event) {
		message = $(this).data('message');
		url = $(this).data('url');
		bootbox.prompt(message, function(result) {
			if (result === null) {
				;
			} else {
				location = url + result;
			}
		});

		event.preventDefault();
	});

	// Gestion RACCOURCIS
	$(document).on('click', '.act-jquery-raccourcis', function(event) {
		$.ajax({url: $(this).data('url') 
			+ '/' + $(this).data('action') 
			+ '/' + $(this).data('iduser')
			+ '/' + $(this).data('idportail')   
			+ '/' + $(this).data('appdico')
			+ '/' + $(this).data('tabledico')
			, success:function(result)
			{ 
				location.reload(); 
			}
		});
		event.preventDefault();
	});

	// CONFIRMATION exec de l'URL
	$(document).on('click', '.act-jquery-confirm', function(event) {
		url = $(this).data('url');
		bootbox.confirm(url, function(result) {
			if (result == true) {
				window.open(url);
			} else {
				;
			}
		});

		event.preventDefault();
	});

	// JQUERY ACTEUR
	// Fonction générale d'appel du controleur - reload de la page en retour
	$(document).on('click', '.act-jquery-reload', function(event) {
		$.ajax({url:$(this).data('url'),success:function(result)
			{ location.reload(); }
		});
		event.preventDefault();
	});

	//VUE iframe=non : Actualisation de la vue dans la portlet
	$(document).on('click', '.act-jquery-portlet-get', function (event) {
		event.preventDefault();
		Metronic.blockUI({target: element, textOnly: true, message: null});
		var element = $(this).closest('.act-jquery-portlet-load');
		var url = $(this).data('url');
		var onChange = element.data('onchange');
		var form = $(this).closest("form");
		$.ajax({
			url: url
			,type: "GET"
				,success:function(result)
				{
					Metronic.unblockUI();
					if ( onChange ) 
					{
						// Remplissage du champ hidden WH_BOUTON
						$('#WH_BOUTON').attr('value', onChange);
						// Validation du formulaire par activation de la fonction submit de la form courante
						form.submit();
					}
					else
					{
						element.click();
					} 
				}
		,error:function(result)
		{  
			Metronic.unblockUI();
			if ( onChange ) 
			{
				// Remplissage du champ hidden WH_BOUTON
				$('#WH_BOUTON').attr('value', onChange);
				// Validation du formulaire par activation de la fonction submit de la form courante
				form.submit();
			}
			else
			{
				element.click();
			} 
		}
		});
	});


	// Chargement du contenu des portlet en AJAX
	$(document).on('click', '.act-jquery-portlet-load', function(event) {
		//var element = $(this).closest(".portlet").children(".portlet-body");
		var $this = $(this);
		var element = $(this);
		var url = $(this).attr("data-url");
		//Metronic.blockUI({target: element, textOnly: true, message: null});
		$.ajax({
			url: url
			,type: "GET"
				,cache: false
				,dataType: "html"
					,success:function(result)
					{
						//Metronic.unblockUI();
						element.empty();
						element.append(result);
					}
		,error:function(result)
		{ 
			//Metronic.unblockUI(); 
			element.html(result); 
		}
		});

		event.preventDefault();
	});

	// Submit d'une portlet en AJAX
	$(document).on('submit', '.act-jquery-portlet-submit', function(event) {
		var element = $(this).closest(".portlet-body");
		Metronic.blockUI({target: element, textOnly: true, message: null});
		$.ajax({
			url: $(this).attr('action')
			,type: $(this).attr('method')
			,cache: false
			,data: $(this).serialize()
			,dataType: "html"
			,success:function(result)
			{
				Metronic.unblockUI();
				element.html(result); 
			}
		,error:function(result)
		{ 
			Metronic.unblockUI(); 
			element.html(result); 
		}
		});
		event.preventDefault();
	});

	// Click sur bouton application
	$('.act-jquery-application').on('show.bs.dropdown', function (event) {
		$(this).children('ul').load($(this).data('url'));
		event.preventDefault();
	});

	// Appel URL dans la même fenêtre
	$(document).on('click', '.act-jquery-url', function (event) {
		var element = $(this).closest(".page-container");
		Metronic.blockUI({target: element, textOnly: true, message: null});
		var $target = $(this).data('target');
		if ( ! $target ||  $target == '' ) {
			window.location = $(this).data('url');
		} else {
			window.open($(this).data('url'), $target);
		}
		event.preventDefault();
		Metronic.unblockUI(element);
	});

	// LISTE avec plugin select2
	$('.act-select2-liste').select2({
		placeholder: ''
		//,formatSelectionCssClass: function (data, container) { return "act-bloc-orange"; }
		,minimumInputLength: 0
		,allowClear: true
	});

	// SUBMIT D'UNE FORM dans une FORMULAIRE ou VUE
	// Message de bloquage de l'interface pour éviter les clics d'impatience de l'utilisateur
	$(document).on('submit', '#WF_MAFORM', function (event)
	{
		var ret = validationFormulaire();
		if ( ret == false || $retour == false ) 
		{
			return false;
		}
		else
		{
			$('button').attr('disabled', 'disabled');
			Metronic.blockUI({target: $('body'), textOnly: true, message: null});
			return true;
		}
		return false;
	}
	);
	
	// Tri des colonnes du tableur
	$(document).on('click', 'th.sort', function(event) {
		$('#WH_SORT').attr('value', $(this).attr('id'));
		$('#WH_SORT_SENS').attr('value', $(this).attr('sortsens'));
		//alert($('#FW_SEARCH').attr('class'));
		$('#WB_SEARCH').click();
	  	event.preventDefault();	    
	});
	// menu sur filtre
	$(document).on('click', '.act-jquery-filtre-menu', function(event) {
		var $value = $(this).data('value');
		var $input = $(this).closest('.input-group').find('input');
		if ( $value == '' )
			$input.attr('value', $value);
		else
			$input.attr('value', '='+$value);
		$('#WB_SEARCH').click();
	  	event.preventDefault();	    
	});

	//Traitement du sélecteur dans la vue
	$(document).on('click', '.act-jquery-selecteur', function(event) {
		$('#WH_BOUTON',window.parent.document).attr('value', $(this).data('onchange'));
		$('#'+$(this).data('set'),window.parent.document).attr('value', $(this).data('get'));
		$('#WF_MAFORM',window.parent.document).submit();
	  	event.preventDefault();	    
	});
	// VueSelecteur
	$(document).on('click', '.act-jquery-vue-selecteur', function(event) {
		var $onchange = $(this).data('onchange');
		var $get_value = $(this).data('get-value');
		var $display_value = $(this).data('display-value');
		var $set_id = $(this).data('set');
		
		$('#' + $set_id,window.opener.document).attr('value', $get_value);
		$('#RF_VUE_SELECTEUR_' + $set_id,window.opener.document).html($display_value);
		if ( $onchange != '' )
		{
			$('#WH_BOUTON',window.opener.document).attr('value', $onchange);
			$('#WF_MAFORM',window.opener.document).submit();
		}
		window.close();
	 	event.preventDefault();	    
	});
	
	/*
	 * Sélection automatique du contenu dans un champ texte 
	 */
	//$(document).on('click','input[type=text]',function(){ this.select(); });
	
	//COLONNES EDITABLES
	// http://jirka.edgering.org/?p=755
	$(".act-jquery-editable").editable( {
		placeholder: ' '
		,emptytext: ''
		,source: function() {
			var id = $(this).data('name'); 
			return RV[id];
		}
		,params: function(params) {
			// Modification des champs qui seront POSTés
			// ajout de WH_ACTION
			params.WH_ACTION = 'UPDATE';
			params.WH_EDITABLE = params.name;
			// le paramètre "name" correspond à l'id du le rubrique (champ) modifiée
			// le paramètre "value" = la valeur champ modifié
			params[params.name] = params.value;
			// On supprimme les paramètres qui risquent de polluer le traitement iter.formulaire()
			delete params.name;
			delete params.value;
			delete params.pk;
			return params;
		} // end params
		,success: function(response, newValue) {
			var $actualiser = $(this).data('actualiser');
			if ( $actualiser && $actualiser == 'oui' )
			{
				actualiser();
			}
		}
	}); // end editable

	$(".act-jquery-editable-entier").editable( {
		placeholder: ' '
		,emptytext: ''
		,params: function(params) {
			// Modification des champs qui seront POSTés
			// ajout de WH_ACTION
			params.WH_ACTION = 'UPDATE';
			params.WH_EDITABLE = params.name;
			// le paramètre "name" correspond à l'id du le rubrique (champ) modifiée
			// le paramètre "value" = la valeur champ modifié
			params[params.name] = params.value;
			// On supprimme les paramètres qui risquent de polluer le traitement iter.formulaire()
			delete params.name;
			delete params.value;
			delete params.pk;
			return params;
		} // end params
		,validate: function(value) {
			var regexp = new RegExp("^[0-9]+$");
			if (value != '' && !regexp.test(value)) {
				return 'Saisie non valide';
			}
		}
		,success: function(response, newValue) {
			var $actualiser = $(this).data('actualiser');
			if ( $actualiser && $actualiser == 'oui' )
			{
				actualiser();
			}
		}
	}); // end editable

	$(".act-jquery-editable-montant").editable( {
		placeholder: ' '
		,emptytext: ''
		,params: function(params) {
			// Modification des champs qui seront POSTés
			// ajout de WH_ACTION
			params.WH_ACTION = 'UPDATE';
			params.WH_EDITABLE = params.name;
			// le paramètre "name" correspond à l'id du le rubrique (champ) modifiée
			// le paramètre "value" = la valeur champ modifié
			params[params.name] = params.value;
			// On supprimme les paramètres qui risquent de polluer le traitement iter.formulaire()
			delete params.name;
			delete params.value;
			delete params.pk;
			return params;
		} // end params
		,display: function(value) {
			$(this).text(value + ' €');
		} 
		,validate: function(value) {
			var regexp = new RegExp("^[0-9]{1,}(\.|)[0-9]{0,2}$");
			if (value != '' && !regexp.test(value)) {
				return 'Saisie non valide';
			}
		}
		,success: function(response, newValue) {
			var $actualiser = $(this).data('actualiser');
			if ( $actualiser && $actualiser == 'oui' )
			{
				actualiser();
			}
		}
	}); // end editable

	$(".act-jquery-editable-taux").editable( {
		placeholder: ' '
		,emptytext: ''
		,params: function(params) {
			// Modification des champs qui seront POSTés
			// ajout de WH_ACTION
			params.WH_ACTION = 'UPDATE';
			params.WH_EDITABLE = params.name;
			// le paramètre "name" correspond à l'id du le rubrique (champ) modifiée
			// le paramètre "value" = la valeur champ modifié
			params[params.name] = params.value;
			// On supprimme les paramètres qui risquent de polluer le traitement iter.formulaire()
			delete params.name;
			delete params.value;
			delete params.pk;
			return params;
		} // end params
		,display: function(value) {
			$(this).text(value + ' %');
		} 
		,validate: function(value) {
			//  (^(100(?:\.0{1,2})?))|(?!^0*$)(?!^0*\.0*$)^\d{1,2}(\.\d{1,2})?$
			//var regexp = new RegExp("^[-]{0,1}[0-9]{0,2}$");
			var regexp = new RegExp("^[0-9]{1,}[0-9]{0,2}$");
			if (value != '' && !regexp.test(value)) {
				return 'Saisie non valide';
			}
		}
		,success: function(response, newValue) {
			var $actualiser = $(this).data('actualiser');
			if ( $actualiser && $actualiser == 'oui' )
			{
				actualiser();
			}
		}
	}); // end editable
	$(".act-jquery-editable-taux2").editable( {
		placeholder: ' '
		,emptytext: ''
		,params: function(params) {
			// Modification des champs qui seront POSTés
			// ajout de WH_ACTION
			params.WH_ACTION = 'UPDATE';
			params.WH_EDITABLE = params.name;
			// le paramètre "name" correspond à l'id du le rubrique (champ) modifiée
			// le paramètre "value" = la valeur champ modifié
			params[params.name] = params.value;
			// On supprimme les paramètres qui risquent de polluer le traitement iter.formulaire()
			delete params.name;
			delete params.value;
			delete params.pk;
			return params;
		} // end params
		,display: function(value) {
			$(this).text(value + ' %');
		} 
		,validate: function(value) {
			var regexp = new RegExp("^[0-9]{1,}(\.|)[0-9]{0,2}$");
			if (value != '' && !regexp.test(value)) {
				return 'Saisie non valide';
			}
		}
		,success: function(response, newValue) {
			var $actualiser = $(this).data('actualiser');
			if ( $actualiser && $actualiser == 'oui' )
			{
				actualiser();
			}
		}
	}); // end editable
	$(".act-jquery-editable-taux3").editable( {
		placeholder: ' '
		,emptytext: ''
		,params: function(params) {
			// Modification des champs qui seront POSTés
			// ajout de WH_ACTION
			params.WH_ACTION = 'UPDATE';
			params.WH_EDITABLE = params.name;
			// le paramètre "name" correspond à l'id du le rubrique (champ) modifiée
			// le paramètre "value" = la valeur champ modifié
			params[params.name] = params.value;
			// On supprimme les paramètres qui risquent de polluer le traitement iter.formulaire()
			delete params.name;
			delete params.value;
			delete params.pk;
			return params;
		} // end params
		,display: function(value) {
			$(this).text(value + ' %');
		} 
		,validate: function(value) {
			var regexp = new RegExp("^[0-9]{1,}(\.|)[0-9]{0,3}$");
			if (value != '' && !regexp.test(value)) {
				return 'Saisie non valide';
			}
		}
		,success: function(response, newValue) {
			var $actualiser = $(this).data('actualiser');
			if ( $actualiser && $actualiser == 'oui' )
			{
				actualiser();
			}
		}
	}); // end editable

	$(".act-jquery-editable-date").editable( {
		placeholder: ' '
		,emptytext: ''
		,datepicker: {
			weekStart: 1
			,dateFormat: 'dd/mm/yyyy'
			,language: 'fr' // bootstrap-datepicker.fr.js modifié /datepicker/bdatepicker/
			,daysOfWeekDisabled: "0,6"
			,calendarWeeks: true
			,todayHighlight: true
		}
		,params: function(params) {
			// Modification des champs qui seront POSTés
			// ajout de WH_ACTION
			params.WH_ACTION = 'UPDATE';
			params.WH_EDITABLE = params.name;
			// le paramètre "name" correspond à l'id du le rubrique (champ) modifiée
			// le paramètre "value" = la valeur champ modifié
			params[params.name] = params.value;
			// On supprimme les paramètres qui risquent de polluer le traitement iter.formulaire()
			delete params.name;
			delete params.value;
			delete params.pk;
			return params;
		} // end params
		,success: function(response, newValue) {
			var $actualiser = $(this).data('actualiser');
			if ( $actualiser && $actualiser == 'oui' )
			{
				actualiser();
			}
		}
	}); // end editable
	// COCHE COLONNE EDITABLE
	$(document).on('click', '.act-jquery-coche', function(event) {
		var element = $(this).closest("table");
		Metronic.blockUI({target: element, textOnly: true, message: null});
		var $url = $(this).attr("data-url");
		var $name = $(this).attr("data-name");
		var $value = $(this).attr("checked") ? '1' : '0';
		var $actualiser = $(this).data('actualiser');
		var $data = {};
		$data['WH_ACTION'] = 'UPDATE'; 
		$data['WH_EDITABLE'] = $name; 
		$data[$name] = $value; 
		$.ajax({
			url: $url
			,type: 'POST'
			,data: $data
			,dataType: "html"
			,success:function(result)
			{
				if ( $actualiser && $actualiser == 'oui' )
				{
					actualiser();
				}
				else
				{
					Metronic.unblockUI(element);
				}
			}
			,error:function(xhr, status, error)
			{
				Metronic.unblockUI(element);
				bootbox.alert('<div class="alert alert-danger">' + xhr.responseText + '</div>');
				//actualiser();
			}
		});
	});

	// BOUTON FORMULAIRE EDITABLE
	$(document).on('click', '.act-jquery-bouton-editable', function(event) {
		var element = $(this).closest("table");
		Metronic.blockUI({target: element, textOnly: true, message: null});
		var $url = $(this).attr("data-url");
		var $name = $(this).attr("data-name");
		var $value = '';
		var $actualiser = $(this).data('actualiser');
		var $data = {};
		$data['WH_ACTION'] = 'UPDATE'; 
		$data['WH_EDITABLE'] = $name; 
		$data[$name] = $value; 
		$.ajax({
			url: $url
			,type: 'POST'
			,data: $data
			,dataType: "html"
			,success:function(result)
			{
				if ( $actualiser && $actualiser == 'oui' )
				{
					actualiser();
				}
				else
				{
					Metronic.unblockUI(element);
				}
			}
			,error:function(xhr, status, error)
			{
				Metronic.unblockUI(element);
				bootbox.alert('<div class="alert alert-danger">' + xhr.responseText + '</div>');
				//actualiser();
			}
		});
	});
	
	// RADIO COLONNE EDITABLE
	$(document).on('click', '.act-jquery-radio-editable', function(event) {
		if ( event.target.nodeName != 'INPUT' )
			return;
		var element = $(this).closest("table");
		Metronic.blockUI({target: element, textOnly: true, message: null});
		var $url = $(this).attr("data-url");
		var $name = $(this).attr("data-name");
		var $id = $(this).attr("data-id");
		var $value = $('input[name="' + $id + '"]:checked').val();
		var $actualiser = $(this).data('actualiser');
		var $data = {};
		$data['WH_ACTION'] = 'UPDATE'; 
		$data['WH_EDITABLE'] = $name; 
		$data[$name] = $value; 
		$.ajax({
			url: $url
			,type: 'POST'
			,data: $data
			,dataType: "html"
			,success:function(result)
			{
				if ( $actualiser && $actualiser == 'oui' )
				{
					actualiser();
				}
				else
				{
					Metronic.unblockUI(element);
				}
			}
			,error:function(xhr, status, error)
			{
				Metronic.unblockUI(element);
				bootbox.alert('<div class="alert alert-danger">' + xhr.responseText + '</div>');
				//actualiser();
			}
		});
		event.preventDefault();
	});

	// A PROPOS
	$(document).on('click', '.act-jquery-apropos', function (event) {
		var hauteur = 'max';
		var largeur = 'l';
		var posx = 'droite';
		window.open($(this).data('url')
	        ,'act_aide'
	        ,get_position_taille_fenetre(posx, null, largeur, hauteur, null));
		event.preventDefault();	    
	});

	// aide en ligne
	$(document).on('click', '.act-jquery-aide', function (event) {
		var hauteur = 'max';
		var largeur = $(this).attr("data-largeur") ? $(this).attr("data-largeur") : 'l';
		var posx = $(this).attr("data-posx") ? $(this).attr("data-posx") : 'gauche';
		var posy = $(this).attr("data-posy") ? $(this).attr("data-posy") : '3';
		var target = $(this).attr("target") ? $(this).attr("target") : 'act-aide';
		window.open($(this).data('url')
	    	    ,target
		        ,get_position_taille_fenetre(posx, posy, largeur, hauteur, null));
		  	event.preventDefault();	    
	});
	
	// aide Tooltip
	$('.act-tooltip').tooltip({
		title: getDataTooltip,
		html: true,
		container: 'body',
		trigger: 'hover',
		placement: 'top'
	});
	// GeData en AJAX 
	function getDataTooltip() {
	    var elem = $(this);
	    var localData = "error";

	    $.ajax(elem.data('url'), {
	        async: false,
	        success: function(data){
	            localData = data;
	        }
	    });
	    return localData;
	}

	$(document).on('click', '.act-jquery-portlet-plierdeplier', function (event) {
		$(this).closest('.portlet-title').children('.tools').find('a').click();
	});
	
})(jQuery);
