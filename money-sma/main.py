import glfw
import OpenGL.GL as gl

import imgui
from imgui.integrations.glfw import GlfwRenderer
from imgui_datascience import imgui_fig as imgui_fig
from matplotlib import pyplot as plt
import numpy as np
from backtesting.test import GOOG

# GOOG.tail()
import pandas as pd


def SMA(values, n):
    """
    Return simple moving average of `values`, at
    each step taking into account `n` previous values.
    """
    return pd.Series(values).rolling(n).mean()

from backtesting import Strategy
from backtesting.lib import crossover
from backtesting import Backtest

import traceback
import logging

from hyperopt import fmin, tpe, rand, hp, Trials, STATUS_OK, STATUS_FAIL,space_eval, trials_from_docs
from hyperopt.exceptions import AllTrialsFailed

def SMABuilder(params):
    n1 = int(params["n1"])
    n2 = int(params["n2"])
    class SmaCross(Strategy):
        def init(self):
            # Precompute the two moving averages
            self.sma1 = self.I(SMA, self.data.Close, n1)
            self.sma2 = self.I(SMA, self.data.Close, n2)
        
        def next(self):
            # If sma1 crosses above sma2, close any existing
            # short trades, and buy the asset
            if crossover(self.sma1, self.sma2):
                self.position.close()
                self.buy()

            # Else, if sma1 crosses below sma2, close any existing
            # long trades, and sell the asset
            elif crossover(self.sma2, self.sma1):
                self.position.close()
                self.sell()
    return SmaCross

def optimizeStrat(strat,market):
    def stratSpace():
        # define a search space
        return hp.choice('a',[
                {
                    "n1":hp.uniform('n1', 10, 1000),
                    "n2":hp.uniform('n2', 10, 1000),
                }
            ])
    def stratbacktest(strat):
        def wrapper(params):
            # params = {"n1":10,"n2":20}
            bt = Backtest(market, strat(params), cash=10000, commission=.002)
            return -bt.run()["Equity Final [$]"]/100
        return wrapper
    func = stratbacktest(strat)
    space = stratSpace()
    trials = Trials()
    trials_step = 1  # how many additional trials to do after loading saved trials. 1 = save after iteration
    best = ""
    # try:
    best = fmin(fn=func, space=space, algo=rand.suggest, trials=trials, max_evals=1000)
    # except Exception as e:
    #     logging.error(traceback.format_exc())
    # try:
    #     tmp = getBestModelfromTrials(trials.trials)
    #     best = tmp
    # except Exception as e:
    #     logging.error(traceback.format_exc())
    # print("Best:", best, stratbacktestSMA(best))
    return best

    # with open("trials"+mkt_name+".json",'a') as f:
    #     json.dump(best,f)
    # return best


# print('Best parameters:')
# print(space_eval(hp_space_reg, best_reg))


# In[9]:
# print(stratbacktest(test))
import crossfiledialog
import json
from numpyencoder import NumpyEncoder

best = {"n1":10,"n2":20}
MKT = GOOG
def optimize():
    global best,MKT
    imgui.begin("Optimize", True)
    choose_market = imgui.button("Choose Market OHLCV .csv")
    if choose_market:
        filename = crossfiledialog.open_file()
        MKT = pd.read_csv(filename)
    optim = imgui.button("Optimize")
    save = imgui.button("Save")
    load = imgui.button("Load")
    if optim:
        best = optimizeStrat(SMABuilder,MKT)
        imgui.text("parameters: "+str(best))
        bt = Backtest(MKT, SMABuilder(best), cash=10000, commission=.002)
        stats = bt.run()
        bt.plot()
        imgui.text("parameters: "+str(best))
        imgui.text_colored(str(stats), 0.2, 1., 0.)
    else:
        bt = Backtest(MKT, SMABuilder(best), cash=10000, commission=.002)
        stats = bt.run()
        imgui.text("parameters: "+str(best))
        imgui.text_colored(str(stats), 0.2, 1., 0.)
    if save:
        save_filename = crossfiledialog.save_file()
        with open(save_filename,"w") as f:
            json.dump(best, f, cls=NumpyEncoder)
    #     imgui.show_demo_window()
    if load:
        filename = crossfiledialog.open_file()
        with open(filename,"r") as f:
            best = json.loads(f.read())
    # imgui.show_demo_window()
    # figure = plt.figure()
    # x = np.arange(0.1, 100, 0.1)
    # y = np.sin(x) / x
    # plt.plot(x, y)
    # imgui_fig.fig(figure, height=700, title="f(x) = sin(x) / x")

    imgui.end()

def main():
    imgui.create_context()
    window = impl_glfw_init()
    impl = GlfwRenderer(window)
    clicked_trade = False
    clicked_test = False
    clicked_optimize = False
    while not glfw.window_should_close(window):
        glfw.poll_events()
        impl.process_inputs()

        imgui.new_frame()

        if imgui.begin_main_menu_bar():
            if imgui.begin_menu("Actions", True):
                clicked_trade, selected_trade = imgui.menu_item(
                    "Trade", 'Cmd+T', False, True
                )
                clicked_test, selected_test = imgui.menu_item(
                    "Test", 'Cmd+E', False, True
                )
                clicked_optimize, selected_optimize = imgui.menu_item(
                    "Optimize", 'Cmd+O', False, True
                )
                clicked_quit, selected_quit = imgui.menu_item(
                    "Quit", 'Cmd+Q', False, True
                )
                if clicked_quit:
                    exit(1)
                imgui.end_menu()
            imgui.end_main_menu_bar()
        
        # if clicked_trade:
        #     trade(True)
        if clicked_optimize:
            optimize()
        
        gl.glClearColor(1., 1., 1., 1)
        gl.glClear(gl.GL_COLOR_BUFFER_BIT)

        imgui.render()
        impl.render(imgui.get_draw_data())
        glfw.swap_buffers(window)

    impl.shutdown()
    glfw.terminate()


def impl_glfw_init():
    width, height = 1280, 720
    window_name = "minimal ImGui/GLFW3 example"

    if not glfw.init():
        print("Could not initialize OpenGL context")
        exit(1)

    # OS X supports only forward-compatible core profiles from 3.2
    glfw.window_hint(glfw.CONTEXT_VERSION_MAJOR, 3)
    glfw.window_hint(glfw.CONTEXT_VERSION_MINOR, 3)
    glfw.window_hint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)

    glfw.window_hint(glfw.OPENGL_FORWARD_COMPAT, gl.GL_TRUE)

    # Create a windowed mode window and its OpenGL context
    window = glfw.create_window(
        int(width), int(height), window_name, None, None
    )
    glfw.make_context_current(window)

    if not window:
        glfw.terminate()
        print("Could not initialize Window")
        exit(1)

    return window


if __name__ == "__main__":
    main()
