// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {Action, Store} from 'redux';
import {debounce} from 'debounce';

import {GlobalState} from 'mattermost-redux/types/store';
import {getTheme} from 'mattermost-redux/selectors/entities/preferences';
import {PluginRegistry} from 'mattermost-webapp/plugins/registry';

import {pluginId} from './manifest';
import {registerCssVars, cleanupCss, isMobile} from './utils/utils';

import IncidentIcon from './components/assets/icons/incident_icon';
import RightHandSidebar from './components/rhs';
import StartIncidentPostMenu from './components/post_menu';
import BackstageModal from './components/backstage/backstage_modal';

import {Hooks} from './hooks';
import {setToggleRHSAction, setBackstageModal} from './actions';
import reducer from './reducer';
import {BackstageArea} from './types/backstage';
import {
    handleWebsocketIncidentUpdate,
    handleWebsocketIncidentCreated,
    handleWebsocketPlaybookCreateModify,
    handleWebsocketPlaybookDelete,
} from './websocket_events';
import {
    WEBSOCKET_INCIDENT_UPDATED,
    WEBSOCKET_INCIDENT_CREATED,
    WEBSOCKET_PLAYBOOK_DELETED,
    WEBSOCKET_PLAYBOOK_CREATED,
    WEBSOCKET_PLAYBOOK_UPDATED,
} from './types/websocket_events';

export default class Plugin {
    public initialize(registry: PluginRegistry, store: Store<object, Action<any>>): void {
        // Ideally, we'd only need to do this within uninitialize, but plugin uninitialization is
        // not called consistently, likely due to https://mattermost.atlassian.net/browse/MM-17087.
        // We rely on the BUILD_TIMESTAMP as exported by webpack to avoid removing our own styles,
        // some of which webpack has already preflighted by the time this file loads.
        cleanupCss(pluginId, BUILD_TIMESTAMP);

        registry.registerReducer(reducer);

        this.updateTheme(store.getState());

        let mainMenuActionId;
        const updateMainMenuAction = () => {
            if (mainMenuActionId && isMobile()) {
                registry.unregisterComponent(mainMenuActionId);
                mainMenuActionId = null;
            } else if (!mainMenuActionId && !isMobile()) {
                mainMenuActionId = registry.registerMainMenuAction(
                    'Incidents & Playbooks',
                    (): void => store.dispatch(setBackstageModal(true, BackstageArea.Incidents)),
                );
            }
        };

        updateMainMenuAction();

        // Would rather use a saga and listen for ActionTypes.UPDATE_MOBILE_VIEW.
        window.addEventListener('resize', debounce(updateMainMenuAction, 300));

        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(RightHandSidebar, null);
        const boundToggleRHSAction = (): void => store.dispatch(toggleRHSPlugin);

        // Store the toggleRHS action to use later
        store.dispatch(setToggleRHSAction(boundToggleRHSAction));

        registry.registerChannelHeaderButtonAction(IncidentIcon, boundToggleRHSAction, 'Incidents', 'Incidents');
        registry.registerPostDropdownMenuComponent(StartIncidentPostMenu);

        registry.registerWebSocketEventHandler(WEBSOCKET_INCIDENT_UPDATED,
            handleWebsocketIncidentUpdate(store.dispatch, store.getState));

        registry.registerWebSocketEventHandler(WEBSOCKET_INCIDENT_CREATED,
            handleWebsocketIncidentCreated(store.dispatch, store.getState));

        registry.registerWebSocketEventHandler(WEBSOCKET_PLAYBOOK_CREATED,
            handleWebsocketPlaybookCreateModify(store.dispatch));

        registry.registerWebSocketEventHandler(WEBSOCKET_PLAYBOOK_UPDATED,
            handleWebsocketPlaybookCreateModify(store.dispatch));

        registry.registerWebSocketEventHandler(WEBSOCKET_PLAYBOOK_DELETED,
            handleWebsocketPlaybookDelete(store.dispatch));

        // Listen to when the theme is loaded
        registry.registerWebSocketEventHandler('preferences_changed',
            () => this.updateTheme(store.getState()));

        const hooks = new Hooks(store);
        registry.registerSlashCommandWillBePostedHook(hooks.slashCommandWillBePostedHook);

        registry.registerRootComponent(BackstageModal);
    }

    public updateTheme(state: GlobalState) {
        const theme = getTheme(state);
        registerCssVars(theme);
    }

    public uninitialize(): void {
        cleanupCss();
    }
}

// @ts-ignore
window.registerPlugin(pluginId, new Plugin());
